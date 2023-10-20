package util

import (
	"context"
	"log"

	"os-container-project/config"
	"os-container-project/model"

	"github.com/segmentio/kafka-go"
)

func KafkaConn(config config.KafkaConfig) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.URL, config.Topic, 0)
	if err != nil {
		log.Fatal("Error connecting to kafka:", err.Error())
	}

	return conn
}

func GenerateKafkaMessages(members []model.Member) []kafka.Message {
	messages := make([]kafka.Message, 0)
	for index := range members {
		messages = append(messages, kafka.Message{
			Key:   []byte(members[index].ID),
			Value: ToJSONBytes(members[index]),
		})
	}
	return messages
}

func IsTopicAlreadyExists(conn *kafka.Conn, topic string) bool {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Fatal("Error reading partitions:", err.Error())
	}

	for _, partition := range partitions {
		if partition.Topic == topic {
			return true
		}
	}

	return false
}
