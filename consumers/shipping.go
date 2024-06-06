package main

import (
	"encoding/json"
	"log"

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
		"shipping", // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
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
			log.Printf("Received a shipping request: %+v", order)

			// Simulate shipping
			d.Ack(false) // Confirm that the message has been processed
			log.Printf("Shipped order: %+v", order)
		}
	}()

	log.Printf(" [*] Waiting for shipping messages. To exit press CTRL+C")
	<-forever
}
