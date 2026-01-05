package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDRIVER   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func LoadConfig() {
	// Try to load .env but don't fail if it's missing — allow real env vars.
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found; falling back to environment variables")
	}

	// read values from environment (possibly set by .env)
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "postgres" // sensible default
	}

	AppConfig = &Config{
		DBDRIVER:   driver,
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}

	fmt.Println(AppConfig)

	log.Println("Config loaded successfully")
}
