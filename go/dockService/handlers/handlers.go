package handlers

import (
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
	rmq, err := rabbitmq.NewRabbitMQ(rmqConn, "dock", "unlock", "unlock")

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
	rmq.Close()

}
