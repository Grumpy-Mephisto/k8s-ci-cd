package main

import (
	"os-container-project/internal/api/routes"
	"os-container-project/internal/config"
	"os-container-project/internal/data"
	"os-container-project/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	app.Static("/", "./public")

	esConfig := config.NewElasicSearchConfig()

	esClient, err := esConfig.NewElasticsearchClient()
	if err != nil {
		log.Fatalf("Error while creating elasticsearch client: %s", err.Error())
	}
	if err := data.SetDefaultData(esClient); err != nil {
		log.Fatalf("Error while setting default data: %s", err.Error())
	}

	routes.SetupRoutes(app, esConfig)

	if err := app.Listen(utils.GetEnv("PORT", ":3000").(string)); err != nil {
		log.Fatalf("Error while listening: %s", err.Error())
	}
}
