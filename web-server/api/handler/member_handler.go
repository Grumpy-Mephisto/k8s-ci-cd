package handler

import (
	"os-container-project/model"

	"github.com/gofiber/fiber/v2"
)

type MemberAPIHandler struct {
	members []model.Member
}

func NewMemberAPIHandler(members []model.Member) *MemberAPIHandler {
	return &MemberAPIHandler{
		members: members,
	}
}

func (h *MemberAPIHandler) GetMembers(c *fiber.Ctx) error {
	return c.JSON(h.members)
}

func (h *MemberAPIHandler) GetMemberByID(c *fiber.Ctx) error {
	memberID := c.Params("id")

	for _, member := range h.members {
		if member.ID == memberID {
			return c.JSON(member)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Member not found")
}

func (h *MemberAPIHandler) AddMember(c *fiber.Ctx) error {
	member := new(model.Member)

	if err := c.BodyParser(member); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	h.members = append(h.members, *member)

	return c.JSON(member)
}

func (h *MemberAPIHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/members", h.GetMembers)
	app.Get("/api/v1/members/:id", h.GetMemberByID)
	app.Post("/api/v1/members", h.AddMember)
}
