package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func LoadConfig() (Config, error) {

	//godotenv.Load() will read the .env file and set the environment variables accordingly. If there is an error loading the .env file, it will log the error and return a wrapped error indicating that the configuration could not be loaded.
	//os.Getenv() is used to retrieve the values of the environment variables that were set by godotenv.Load(). These values are then assigned to the respective fields in the Config struct, which is returned to the caller.
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return Config{}, fmt.Errorf("could not load configuration: %w", err)
	}
	mongoURI, err := extractenv("MONGO_URI")
	if err != nil {
		return Config{}, fmt.Errorf("could not load configuration: %w", err)
	}
	mongoDB, err := extractenv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, fmt.Errorf("could not load configuration: %w", err)
	}
	serverPort, err := extractenv("PORT")
	if err != nil {
		return Config{}, fmt.Errorf("could not load configuration: %w", err)
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: serverPort,
	}, nil
}

func extractenv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}
