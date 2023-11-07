package main

import (
	"fmt"
	"go-bike-share/dockService/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/dock/unlock", handlers.UnlockDock)

	PORT := os.Getenv("PORT")
	fmt.Println("Listening on port -", PORT)
	err = http.ListenAndServe(":"+PORT, mux)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
