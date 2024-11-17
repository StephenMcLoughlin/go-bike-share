package handlers

import (
	"fmt"
	"go-bike-share/shared/postgres"
	"go-bike-share/shared/rabbitmq"
	"net/http"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func UnlockDock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	dockId, err := strconv.Atoi(r.URL.Query().Get("dockid"))

	if err != nil || dockId < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	rmqConn := os.Getenv("RABBITMQ_URL")
	fmt.Println(rmqConn)
	rmq, err := rabbitmq.NewRabbitMQ("amqp://guest:guest@rabbitmq-mqtt:5672", "dock", "unlock", "unlock")

	fmt.Println(rmq)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	err = rmq.Channel.Publish("dock", "unlock", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(strconv.Itoa(dockId)),
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	pg, err := postgres.NewPostgres("bike-share-db", 5432, "postgres", "postgres", "postgres")
	result, err := pg.ExecQuery("INSERT INTO public.test (dockid) VALUES ($1)", dockId)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	rmq.Close()

}

func ReportDock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
