package config

import (
	"database/sql"
	"fmt"
	"log"
)

type InfraConfig struct {
	*sql.DB
}

func NewInfra(cfg Config) InfraConfig {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.POSTGRES_USER, cfg.POSTGRES_PASSWORD, cfg.DB_NAME)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Database connection established")
	return InfraConfig{
		DB: db,
	}

}
