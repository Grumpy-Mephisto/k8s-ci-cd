package main

import (
	"log"
	"os-container-project/api/handler"
	"os-container-project/api/routes"
	"os-container-project/config"
	"os-container-project/data"
	"os-container-project/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	kafkaConfig := config.LoadKafkaConfig()
	kafkaConn := util.KafkaConn(*kafkaConfig)
	defer kafkaConn.Close()

	members := data.LoadMembersData()

	memberHandler := handler.NewMemberAPIHandler(members)

	routes.SetupRoutes(app, memberHandler)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
