package main

import (
	"os-container-project/internal/api/routes"
	"os-container-project/internal/config"
	"os-container-project/internal/model"
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

	mariaDBConfig := config.NewMariaDBConfig()
	mariadbClient, err := mariaDBConfig.NewMariaDBClient()
	if err != nil {
		log.Fatalf("Error while creating mariadb client: %s", err.Error())
	}
	if err := mariadbClient.AutoMigrate(&model.Member{}); err != nil {
		log.Fatalf("Error while migrating: %s", err.Error())
	}
	if err := utils.SetDefaultData(mariadbClient); err != nil {
		log.Fatalf("Error while setting default data: %s", err.Error())
	}

	redisConfig := config.NewRedisConfig()

	routes.SetupRoutes(app, mariadbClient, redisConfig)

	if err := app.Listen(utils.GetEnv("PORT", ":3000").(string)); err != nil {
		log.Fatalf("Error while listening: %s", err.Error())
	}
}
