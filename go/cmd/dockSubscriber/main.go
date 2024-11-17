package main

import (
	"fmt"
	"go-bike-share/shared/postgres"
	"go-bike-share/shared/rabbitmq"
	"os"
	"sync"
)

type Server struct {
	rmq       *rabbitmq.RabbitMQ
	db        *postgres.Postgres
	done      chan struct{}
	queueName string
}

func main() {

	rabbitmqUrl := os.Getenv("RABBITMQ_URL")
	rmq, err := rabbitmq.NewRabbitMQ(rabbitmqUrl, "dock", "unlocked", "unlocked")

	if err != nil {
		fmt.Println(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := postgres.NewPostgres(dbHost, 5432, dbUser, dbPassword, dbName)

	if err != nil {
		fmt.Println(err.Error())
	}

	svr := &Server{
		rmq:       rmq,
		db:        db,
		done:      make(chan struct{}),
		queueName: "unlocked",
	}

	svr.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
		close(svr.done)
	}
}

func (s *Server) ListenAndServe() error {
	msgs, err := s.rmq.Channel.Consume(
		s.queueName, // queue
		s.queueName, // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func(db *postgres.Postgres) {
		defer db.Close()
		defer wg.Done()

		for {
			select {
			case <-s.done:
				fmt.Println("Done")
			case msg, ok := <-msgs:
				if !ok {
					fmt.Println("Message channel closed")
					return
				}

				message := string(msg.Body)
				fmt.Println(message)
				// msg.Nack(false, true)
			}
		}
	}(s.db)

	wg.Wait()

	return nil

}
