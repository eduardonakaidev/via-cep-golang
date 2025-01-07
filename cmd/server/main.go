package main

import (
	"log"
	"net/http"

	"github.com/eduardofrnkdev/via-cep-golang/internal/handlers"
)

func main() {
	router := handlers.SetupRoutes()

	log.Println("Server running on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
