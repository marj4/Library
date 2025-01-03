package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

const (
	ErrorLoadEnv = "Error loading .env file: %v"
)

func StartConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		DatabaseURL: os.Getenv("DATABASE"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}, nil

}
