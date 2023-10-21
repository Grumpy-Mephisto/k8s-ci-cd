package routes

import (
	"os-container-project/internal/api/handlers"
	"os-container-project/internal/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, esConfig *config.ElasticsearchConfig) {
	v1 := app.Group("/api/v1")

	memberHandler := handlers.NewMemberHandler(esConfig)
	member := v1.Group("/members")
	member.Get("/", memberHandler.GetMembers)
	member.Get("/:id", memberHandler.GetMemberByID)
	member.Post("/", memberHandler.AddMemberHandler)
}
