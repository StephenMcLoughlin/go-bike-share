package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQ(url string, exchangeName string, queueName string, bindingKey string) (*RabbitMQ, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type defaults topic
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	err = channel.QueueBind(
		queueName,    // Queue name
		bindingKey,   // Routing key
		exchangeName, // Exchange name
		false,        // NoWait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind queue to exchange: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return &RabbitMQ{
		Connection: connection,
		Channel:    channel,
	}, nil
}

func (r *RabbitMQ) Close() {
	r.Connection.Close()
}
