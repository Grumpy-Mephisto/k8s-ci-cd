package handlers

import (
	"context"
	"encoding/json"

	"os-container-project/internal/config"
	"os-container-project/internal/model"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	esConfig *config.ElasticsearchConfig
}

func NewMemberHandler(config *config.ElasticsearchConfig) *MemberHandler {
	return &MemberHandler{esConfig: config}
}

func (h *MemberHandler) GetMembers(c *fiber.Ctx) error {
	esClient, err := h.esConfig.NewElasticsearchClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while creating elasticsearch client"})
	}

	result, err := esClient.Search().Index("members").Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying elasticsearch"})
	}

	var members []model.Member
	for _, hit := range result.Hits.Hits {
		var member model.Member
		if err := json.Unmarshal(hit.Source, &member); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error processing member data"})
		}
		members = append(members, member)
	}

	return c.JSON(members)
}

func (h *MemberHandler) GetMemberByID(c *fiber.Ctx) error {
	esClient, err := h.esConfig.NewElasticsearchClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while creating elasticsearch client"})
	}

	memberID := c.Params("id")

	result, err := esClient.Get().Index("members").Id(memberID).Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying elasticsearch"})
	}

	if !result.Found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Member not found"})
	}

	var member model.Member

	if err := json.Unmarshal(result.Source, &member); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error processing member data"})
	}

	return c.JSON(member)
}

func (h *MemberHandler) AddMemberHandler(c *fiber.Ctx) error {
	esClient, err := h.esConfig.NewElasticsearchClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while creating elasticsearch client"})
	}

	var memberData model.Member
	if err := c.BodyParser(&memberData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON"})
	}

	_, err = esClient.Index().Index("members").BodyJson(memberData).Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while indexing"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Member added successfully"})
}
