package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error

	if DB, err = gorm.Open(postgres.Open(config.PostgresURL), &gorm.DB{}); err != nil {
		log.Fatal("Failed to connect to Database!")
	}
	fmt.Println("Connected successfully to Database")
}
