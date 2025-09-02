package server

import (
	"lms-backend/internal/config"
	"lms-backend/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; relying on environment variables")
	}

	config.ConnectDB()

	r := routes.SetupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
}
