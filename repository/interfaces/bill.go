package interfaces

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
	"github.com/jackc/pgx/v5"
)

type BillRepository interface {
	// produce bill
	CreateBill(ctx context.Context, db *pgx.Conn, payload entities.TblBill) (entities.TblBill, error)
}
