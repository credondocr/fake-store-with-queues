package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"

	"github.com/credondocr/fake-store-with-queues/common"
)

func main() {
	conn, ch := common.GetChannel()
	defer conn.Close()
	defer ch.Close()

	queues := []string{"order_validation", "payment_processing", "shipping"}
	for _, queue := range queues {
		_, err := ch.QueueDeclare(
			queue, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		common.FailOnError(err, "Failed to declare a queue")
	}

	for i := 0; i < 10; i++ {
		order := common.Order{
			ID:         i,
			Product:    "Product " + fmt.Sprint(i),
			Quantity:   1,
			Price:      10.0 + float64(i),
			CustomerID: 12345 + i,
		}
		body, err := json.Marshal(order)
		common.FailOnError(err, "Failed to marshal order to JSON")
		err = ch.Publish(
			"",                 // exchange
			"order_validation", // routing key
			false,              // mandatory
			false,              // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		common.FailOnError(err, "Failed to publish a message")
		log.Printf("Sent %s", body)
		time.Sleep(1 * time.Second)
	}
}
