package interfaces

import (
	"context"
	"database/sql"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

type BillRepository interface {
	// produce bill
	CreateBill(ctx context.Context, db *sql.DB, payload entities.TblBill) (entities.TblBill, error)
}
