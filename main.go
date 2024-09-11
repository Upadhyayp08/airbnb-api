package main

import (
	"airbnb-api/database"
	router "airbnb-api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to MongoDB
    database.ConnectDB()

    // Initialize Router
    r := router.InitializeRouter()

    // Start the server
    fmt.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
