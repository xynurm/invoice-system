package main

import (
	"fmt"
	"invoice-system/migrations"
	"invoice-system/pkg/mysql"
	"invoice-system/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load configuration environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load godotenv")
	}

	// Initialize DB connection
	mysql.DatabaseInit()

	// Run database migrations
	migrations.RunMigration()

	// Initialize Gin router
	r := gin.Default()

	// Initialize and configure routes
	routes.RouteInit(r.Group("/api/v1"))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT .env is required...")
	}

	// Run server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	fmt.Println("Server running on port " + port + "...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
