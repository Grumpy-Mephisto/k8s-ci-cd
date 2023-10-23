package routes

import (
	"os-container-project/api/handlers"
	"os-container-project/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, redisClient *config.RedisConfig) {
	v1 := app.Group("/api/v1")

	memberHandler := handlers.NewMemberHandler(db, redisClient)
	member := v1.Group("/members")
	member.Get("/", memberHandler.GetMembers)
	member.Get("/:student_id", memberHandler.GetMemberByID)
	member.Post("/", memberHandler.AddMemberHandler)

	postHandler := handlers.NewPostHandler(db, redisClient)
	post := v1.Group("/posts")
	post.Get("/", postHandler.GetPosts)
	post.Post("/", postHandler.AddPost)
}
