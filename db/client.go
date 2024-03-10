package db

import (
	"fmt"
	"log"

	"github.com/k0kishima/golang-realworld-example-app/config"
)

func GetDataSourceName() string {
	// Environment variables are already loaded when the application is launched.
	dbConfig, err := config.GetDBConfig()
	if err != nil {
		log.Fatalf("Error getting database configuration: %v", err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)
}
