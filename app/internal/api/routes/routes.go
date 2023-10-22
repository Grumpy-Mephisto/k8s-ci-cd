package routes

import (
	"os-container-project/internal/api/handlers"
	"os-container-project/internal/config"

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
}
