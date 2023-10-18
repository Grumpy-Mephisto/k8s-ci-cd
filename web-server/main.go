package main

import (
	"log"
	"os-container-project/api/handler"
	"os-container-project/config"
	"os-container-project/model"
	"os-container-project/pkg/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       config.GetEnvAsBool("PREFORK", false),
		ServerHeader:  config.GetEnv("SERVER_HEADER", "Fiber"),
		StrictRouting: config.GetEnvAsBool("STRICT_ROUTING", false),
		Concurrency:   config.GetEnvAsInt("CONCURRENCY", 256),
		BodyLimit:     config.GetEnvAsInt("BODY_LIMIT", 4*1024*1024),
	})

	kafkaURL := config.GetEnv("KAFKA_URL", "localhost:9092")
	kafkaTopic := config.GetEnv("KAFKA_TOPIC", "os-container-project")

	kafkaConfig := config.KafkaConfig{
		URL:   kafkaURL,
		Topic: kafkaTopic,
	}

	conn := util.KafkaConn(kafkaConfig)
	defer conn.Close()

	if !util.IsTopicAlreadyExists(conn, kafkaConfig.Topic) {
		topicConfig := []kafka.TopicConfig{
			{
				Topic:             kafkaConfig.Topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfig...)
		if err != nil {
			log.Fatal("Error creating topic:", err)
		}
	}

	var members []model.Member
	if err := util.LoadJSONFile("data/members.json", &members); err != nil {
		log.Fatal("Error loading members from JSON:", err)
	}

	messages := util.GenerateKafkaMessages(members)

	if err := conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal("Failed to set read deadline:", err)
	}
	_, err := conn.WriteMessages(messages...)
	if err != nil {
		log.Fatal("Failed to write messages:", err)
	}

	memberAPIHandler := handler.NewMemberAPIHandler(members)
	memberAPIHandler.RegisterRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
