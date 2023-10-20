package routes

import (
	"os-container-project/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, memberHandler *handler.MemberAPIHandler) {
	v1 := app.Group("/api/v1")

	member := v1.Group("/members")
	member.Get("/", memberHandler.GetMembers)
	member.Get("/:id", memberHandler.GetMemberByID)
	member.Post("/", memberHandler.AddMember)
}
