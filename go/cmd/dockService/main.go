package main

import (
	"fmt"
	"go-bike-share/cmd/dockService/handlers"
	"net/http"
	"os"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	mux := http.NewServeMux()

	mux.HandleFunc("/dock/unlock", handlers.UnlockDock)
	mux.HandleFunc("/dock/report", handlers.ReportDock)

	PORT := os.Getenv("PORT")
	fmt.Println("Listening on port -", PORT)
	err := http.ListenAndServe(":"+PORT, mux)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
