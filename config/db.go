package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

// Database modal
type Database struct {
	*pgx.Conn
}

// NewDatabase creates a new database instance
func NewDatabase(cfg Config) Database {

	conn, err := pgx.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		log.Panic(err)
	}
	return Database{Conn: conn}
}
