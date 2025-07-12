package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUSER                   string
	DBPASS                   string
	DBHOST                   string
	DBPORT                   string
	DBNAME                   string
}

func LoadConfig() *Config {
	// Check if the .env file exists
	if _, err := os.Stat("/.env"); err == nil {
		// Load the .env file
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	return &Config{
		DBUSER:                   GetEnv("DBUSER", "default_user"),
		DBPASS:                   GetEnv("DBPASS", "default_pass"),
		DBHOST:                   GetEnv("DBHOST", "localhost"),
		DBPORT:                   GetEnv("DBPORT", "3306"),
		DBNAME:                   GetEnv("DBNAME", "postalnav"),
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Environment variable %s not found, using default value", key)
	return defaultValue
}
