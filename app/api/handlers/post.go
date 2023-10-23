package handlers

import (
	"os-container-project/config"
	"os-container-project/internal/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type PostHandler struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewPostHandler(db *gorm.DB, redisConfig *config.RedisConfig) *PostHandler {
	return &PostHandler{
		db:          db,
		redisClient: redisConfig.NewRedisClient(),
	}
}

func (h *PostHandler) GetPosts(c *fiber.Ctx) error {
	postsRedis, err := h.getPostsFromRedis()
	if err == nil {
		return c.JSON(postsRedis)
	}

	var posts []model.Post

	if err := h.db.Preload("Author").Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying the database"})
	}

	if err := h.setPostsInRedis(posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error caching data in Redis"})
	}

	return c.JSON(posts)
}

func (h *PostHandler) AddPost(c *fiber.Ctx) error {
	var postData model.Post
	if err := c.BodyParser(&postData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON"})
	}

	var author model.Member
	if err := h.db.First(&author, postData.AuthorID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Author not found"})
	}

	postData.Author = author

	if err := h.db.Create(&postData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while adding a post"})
	}

	if err := h.clearPostsInRedis(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error clearing cached data in Redis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post added successfully"})
}

func (h *PostHandler) getPostsFromRedis() ([]model.Post, error) {
	ctx := context.Background()
	key := "posts"
	val, err := h.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var posts []model.Post
	err = json.Unmarshal([]byte(val), &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (h *PostHandler) setPostsInRedis(posts []model.Post) error {
	ctx := context.Background()
	key := "posts"
	val, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	return h.redisClient.Set(ctx, key, val, 24*time.Hour).Err()
}

func (h *PostHandler) clearPostsInRedis() error {
	ctx := context.Background()
	key := "posts"
	return h.redisClient.Del(ctx, key).Err()
}
