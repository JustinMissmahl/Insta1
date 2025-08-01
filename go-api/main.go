package main

import (
	"log"
	"os"

	"instagram-downloader-api/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := server.New()
	log.Printf("Starting server on port %s", port)
	if err := srv.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
