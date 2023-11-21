package manager

import (
	"context"
	"fmt"
	"log"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
	"github.com/jackc/pgx/v5"
)

type InfraManager interface {
	Conn() *pgx.Conn
}

type infraManager struct {
	conn *pgx.Conn
	cfg  *config.Config
}

func (i *infraManager) openConn() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.POSTGRES_HOST, i.cfg.POSTGRES_PORT, i.cfg.POSTGRES_USER, i.cfg.POSTGRES_PASSWORD, i.cfg.DB_NAME)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Panic(err)
	}
	i.conn = conn
	return nil
}

func (i *infraManager) Conn() *pgx.Conn {
	return i.conn
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}
	return conn, nil
}
