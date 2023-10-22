package handlers

import (
	"context"
	"encoding/json"

	"os-container-project/internal/config"
	"os-container-project/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/olivere/elastic/v7"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while creating Elasticsearch client"})
	}

	memberID := c.Params("id")

	result, err := esClient.Get().Index("members").Id(memberID).Do(context.Background())
	if err != nil {
		if elastic.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Member not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying Elasticsearch"})
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

	exists, err := esClient.Exists().Index("members").Id(memberData.ID).Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error querying elasticsearch"})
	}

	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Member already exists"})
	}

	_, err = esClient.Update().
		Index("members").
		Id(memberData.ID).
		Script(elastic.NewScriptInline("ctx._source = params.data").Lang("painless").Param("data", memberData)).
		Upsert(memberData).
		Do(context.Background())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error while updating Elasticsearch"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Member added successfully"})
}
