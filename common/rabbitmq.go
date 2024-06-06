package common

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func GetChannel() (*amqp.Connection, *amqp.Channel) {
	rootPath, err := filepath.Abs("../")
	if err != nil {
		log.Fatalf("Failed to get root path: %s", err)
	}

	envPath := filepath.Join(rootPath, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		log.Fatalf("RABBITMQ_URL not set in environment")
	}

	conn, err := amqp.Dial(rabbitmqURL)
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return conn, ch
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
