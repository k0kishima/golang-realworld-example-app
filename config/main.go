package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetDBConfig() (*DBConfig, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" {
		return nil, errors.New("DB_USER is not set")
	}
	if dbPassword == "" {
		return nil, errors.New("DB_PASSWORD is not set")
	}
	if dbHost == "" {
		return nil, errors.New("DB_HOST is not set")
	}
	if dbName == "" {
		return nil, errors.New("DB_NAME is not set")
	}

	return &DBConfig{
		User:     dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Name:     dbName,
	}, nil
}
