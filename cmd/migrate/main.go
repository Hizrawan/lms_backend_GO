package main

import (
	"lms-backend/internal/config"
	migrate "lms-backend/migrations"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	migrate.Run()
}
