
# Fake Store with Queues

This project is a simulation of a fake store that processes orders using RabbitMQ. The application demonstrates the use of RabbitMQ for message queuing, with multiple producers and consumers handling different stages of order processing.

## Project Structure

```
fake-store-with-queues
│
├── go.mod
├── producer
│   └── main.go
│
├── consumers
│   ├── validator.go
│   ├── payment_processor.go
│   └── shipping.go
│
└── common
    ├── rabbitmq.go
    └── types.go
```

### Modules and Dependencies

- `producer/main.go`: Simulates order generation and sends orders to the `order_validation` queue.
- `consumers/validator.go`: Consumes messages from the `order_validation` queue, validates them, and forwards them to the `payment_processing` queue.
- `consumers/payment_processor.go`: Consumes messages from the `payment_processing` queue, processes payments, and forwards the messages to the `shipping` queue.
- `consumers/shipping.go`: Consumes messages from the `shipping` queue and simulates order shipping.
- `common/rabbitmq.go`: Contains common RabbitMQ setup and utility functions.
- `common/types.go`: Defines the `Order` struct used throughout the project.

## Prerequisites

- Go 1.16 or higher
- RabbitMQ server running locally (`amqp://guest:guest@localhost:5672/`)
- Try using a tier free Cloudmqp instance !!

## Getting Started

1. **Clone the repository**:

   ```sh
   git clone https://github.com/credondocr/fake-store-with-queues.git
   cd fake-store-with-queues
   ```

2. **Create a .env file**:
Create a .env file in the root of the project with the following content:
   ```sh
   RABBITMQ_URL=amqp://guest:guest@localhost:5672/
   ```

3. **Initialize Go modules**:

   ```sh
   go mod tidy
   ```

4. **Run the producer**:

   ```sh
   go run producer/main.go
   ```

5. **Run the consumers in separate terminals**:

   ```sh
   go run consumers/validator.go
   go run consumers/payment_processor.go
   go run consumers/shipping.go
   ```

## Order Processing Flow

1. The **producer** generates order messages in JSON format and sends them to the `order_validation` queue.
2. The **validator consumer** reads messages from the `order_validation` queue, unmarshals the JSON into an `Order` struct, validates the order, and forwards it to the `payment_processing` queue.
3. The **payment processor consumer** reads messages from the `payment_processing` queue, unmarshals the JSON into an `Order` struct, processes the payment, and forwards the order to the `shipping` queue.
4. The **shipping consumer** reads messages from the `shipping` queue, unmarshals the JSON into an `Order` struct, and simulates the shipping of the order.


## License

This project is licensed under the MIT License.
