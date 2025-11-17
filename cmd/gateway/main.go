package main

import (
	"log"
	"net/http"

	"golite-gateway/internal/config"
	"golite-gateway/internal/router"
)

func main() {
	// Load config
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Start watcher to reload config on file changes
	go config.Watch("config.yaml")

	// Create router with dynamic routes
	mux := router.New(cfg)

	log.Println("API Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
