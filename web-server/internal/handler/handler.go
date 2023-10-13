package handler

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Handle(c *fiber.Ctx) error
}

type (
	HealthHandler   struct{}
	ArgoHandler     struct{}
	MemberHandler   struct{}
	MemberIdHandler struct{}
	MemberData      struct {
		ID       int    `json:"id"`
		Fullname string `json:"fullname"`
	}
)

func (h *HealthHandler) Handle(c *fiber.Ctx) error {
	if err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "API is healthy",
	}); err != nil {
		return err
	}

	return nil
}

func (h *ArgoHandler) Handle(c *fiber.Ctx) error {
	if err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "CI/CD pipeline is working!",
		"success":   true,
		"timestamp": time.Now().Unix(),
	}); err != nil {
		return err
	}

	return nil
}

func (h *MemberHandler) Handle(c *fiber.Ctx) error {
	memberJSON, err := os.ReadFile("data/members.json")
	if err != nil {
		log.Printf("Error reading JSON file: %v", err)
		return err
	}

	var members []MemberData
	if err := json.Unmarshal(memberJSON, &members); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"members": members,
	}); err != nil {
		log.Printf("Error sending JSON response: %v", err)
		return err
	}

	return nil
}

func (h *MemberIdHandler) Handle(c *fiber.Ctx) error {
	memberJSON, err := os.ReadFile("data/members.json")
	if err != nil {
		return err
	}

	var members []MemberData
	if err := json.Unmarshal(memberJSON, &members); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	for _, member := range members {
		if member.ID == id {
			if err := c.Status(fiber.StatusOK).JSON(fiber.Map{
				"member": member,
			}); err != nil {
				return err
			}

			return nil
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}
