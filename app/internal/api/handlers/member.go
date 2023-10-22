package handlers

import (
	"context"
	"encoding/json"
	"os-container-project/internal/config"
	"os-container-project/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MemberHandler struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewMemberHandler(db *gorm.DB, redisClient *config.RedisConfig) *MemberHandler {
	return &MemberHandler{db: db, redisClient: redisClient.NewRedisClient()}
}

func (h *MemberHandler) GetMembers(c *fiber.Ctx) error {
	var members []model.Member
	if err := h.db.Find(&members).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying the database"})
	}
	return c.JSON(members)
}

func (h *MemberHandler) GetMemberByID(c *fiber.Ctx) error {
	studentID := c.Params("student_id")

	member, err := h.getMemberFromRedis(studentID)
	if err == nil {
		return c.JSON(member)
	}

	var memberFromDB model.Member
	if err := h.db.First(&memberFromDB, "student_id = ?", studentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Member not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying the database"})
	}

	if err := h.setMemberInRedis(studentID, &memberFromDB, 24*time.Hour); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error caching data in Redis"})
	}

	return c.JSON(memberFromDB)
}

func (h *MemberHandler) AddMemberHandler(c *fiber.Ctx) error {
	var memberData model.Member
	if err := c.BodyParser(&memberData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON"})
	}

	var count int64
	h.db.Model(&model.Member{}).Where("student_id = ?", memberData.StudentID).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Member already exists"})
	}

	if err := h.db.Create(&memberData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while adding a member"})
	}

	if err := h.setMemberInRedis(memberData.StudentID, &memberData, 24*time.Hour); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error caching data in Redis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Member added successfully"})
}

func (h *MemberHandler) getMemberFromRedis(studentID string) (*model.Member, error) {
	ctx := context.Background()
	key := "member:" + studentID
	val, err := h.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var member model.Member
	err = json.Unmarshal([]byte(val), &member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (h *MemberHandler) setMemberInRedis(studentID string, member *model.Member, expiration time.Duration) error {
	ctx := context.Background()
	key := "member:" + studentID
	val, err := json.Marshal(member)
	if err != nil {
		return err
	}

	return h.redisClient.Set(ctx, key, val, expiration).Err()
}
