package main

import (
	"fmt"
	"sync"
)

type Database struct {
}

type Server struct {
	db  *Database
	rmq *string
}

func NewServer(db *Database, rmq *string) *Server {
	return &Server{
		db:  db,
		rmq: rmq,
	}
}

type Message struct {
	Content string
}

var messageChannel = make(chan Message)

func worker(done <-chan struct{}, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case msg := <-messageChannel:
			fmt.Println(msg.Content)

		case <-done:
			fmt.Printf("Worker %d done", id)
			return
		}
	}
}

func test(done <-chan struct{}) {

	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("test")
		}
	}
}

func main() {

	// var wg sync.WaitGroup

	// done := make(chan struct{})

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go worker(done, &wg, i)
	// }

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		messageChannel <- Message{Content: fmt.Sprintf("Message %d", i)}
	// 	}
	// }()

	// close(done)

	// wg.Wait()

	done := make(chan struct{})

	go test(done)

	<-done
	// for {
	// 	fmt.Println("test")
	// }
}
