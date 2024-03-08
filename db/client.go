package db

import (
	"fmt"
	"os"
)

func GetDataSourceName() string {
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "127.0.0.1:53306")
	dbName := getEnv("DB_NAME", "golang_realworld")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbHost, dbName)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
