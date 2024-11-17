package main

import (
	"fmt"
	"go-bike-share/shared/rabbitmq"
	"log"
)

type Server struct {
	rmq  *rabbitmq.RabbitMQ
	done chan struct{}
}

func main() {
	errC, err := run()
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run() (<-chan error, error) {
	rmq, err := rabbitmq.NewRabbitMQ("amqp://guest:guest@localhost:5672", "tasks", "tasks", "tasks")
	if err != nil {
		fmt.Println(err.Error())
	}

	svr := &Server{
		rmq:  rmq,
		done: make(chan struct{}),
	}

	errC := make(chan error, 1)

	go func() {
		fmt.Println("listening and serving")
		err := svr.ListenAndServe()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return errC, nil
}

func (s *Server) ListenAndServe() error {

	queue, err := s.rmq.Channel.QueueDeclare(
		"tasks", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	err = s.rmq.Channel.QueueBind(
		queue.Name, // queue name
		"tasks",    // routing key
		"tasks",    // exchange
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := s.rmq.Channel.Consume(
		"tasks", // queue
		"tasks", // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)

	go func() {
		for msg := range msgs {
			message := string(msg.Body)
			fmt.Println(message)
			//  s.done <- struct{}{}
		}
	}()
	return nil
}
