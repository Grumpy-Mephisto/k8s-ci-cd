package routes

import (
	"os-container-project/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, memberHandler *handler.MemberAPIHandler) {
	app.Get("/api/v1/members", memberHandler.GetMembers)
	app.Get("/api/v1/members/:id", memberHandler.GetMemberByID)
	app.Post("/api/v1/members", memberHandler.AddMember)
}
