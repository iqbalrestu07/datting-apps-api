package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	DBDriver string
	// psql
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
}

func GetAPPConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		AppPort:    os.Getenv("SERVER_ADDRESS"),
		DBDriver:   os.Getenv("DATABASE_DRIVER"),
		DBHost:     os.Getenv("DATABASE_HOST"),
		DBPort:     os.Getenv("DATABASE_PORT"),
		DBUser:     os.Getenv("DATABASE_USER"),
		DBPassword: os.Getenv("DATABASE_PASSWORD"),
		DBName:     os.Getenv("DATABASE_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
