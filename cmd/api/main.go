package main

import (
	"crud-api/internal/config"
	"crud-api/internal/db"
	"crud-api/internal/server"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig() // Load the configuration settings from the .env file using the LoadConfig function defined in the internal/config/config.go file. This function reads the environment variables and returns a Config struct containing the necessary configuration values.
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the MongoDB database using the configuration settings. The Connect function is defined in the internal/db/mongo.go file and establishes a connection to the MongoDB server, returning a client and a database instance.
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("Error disconnecting from database: %v", err)
		}
	}()

	router := server.NewRouter(database) // Create a new router using the NewRouter function defined in the internal/server/router.go file. This function sets up the HTTP routes and handlers for the API.
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil { // Start the HTTP server on the specified address and port. The Run method of the Gin router will block and listen for incoming HTTP requests until the server is stopped. If there is an error starting the server, it will log the error and exit the application.
		log.Fatalf("Failed to start server: %v", err)
	}

}
