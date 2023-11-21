package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type InfraConfig struct {
	*pgx.Conn
}

func NewInfra(cfg *Config) InfraConfig {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.POSTGRES_USER, cfg.POSTGRES_PASSWORD, cfg.DB_NAME)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Database connection established")
	return InfraConfig{
		Conn: conn,
	}

}
