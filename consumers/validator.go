package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/credondocr/fake-store-with-queues/common"
)

func main() {
	conn, ch := common.GetChannel()
	defer conn.Close()
	defer ch.Close()

	err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	common.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		"order_validation", // queue
		"",                 // consumer
		false,              // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	common.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var order common.Order
			err := json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Printf("Failed to unmarshal order: %s", err)
				d.Nack(false, false)
				continue
			}
			log.Printf("Received a validation request: %+v", order)

			// Simulate validation
			d.Ack(false) // Confirm that the message has been processed

			body, err := json.Marshal(order)
			if err != nil {
				log.Printf("Failed to marshal order to JSON: %s", err)
				continue
			}

			err = ch.Publish(
				"",                   // exchange
				"payment_processing", // routing key
				false,                // mandatory
				false,                // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			common.FailOnError(err, "Failed to publish a message")
			log.Printf("Validated and sent to payment processing: %s", body)
		}
	}()

	log.Printf(" [*] Waiting for validation messages. To exit press CTRL+C")
	<-forever
}
