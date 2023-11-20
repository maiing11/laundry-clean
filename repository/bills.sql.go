package repository

import (
	"context"
	"time"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
	"github.com/jackc/pgx/v5"
)

type BillRepository interface {
	// produce bill
	CreateBill(ctx context.Context, db *pgx.Conn, payload entities.TblBill) (entities.TblBill, error)
}

const createBill = `INSERT INTO tbl_bills (
	date_in, date_finished, customer_id, services, bill_amount
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, date_in, date_finished, customer_id, services, bill_amount
)`

const createServices = `INSERT INTO tbl_services service_no, quantity, subtotal
VALUES $1, $2, $3
RETURNING id, service_no, quantity, subtotal
`

func (q *Queries) CreateBill(ctx context.Context, db *pgx.Conn, payload entities.TblBill) (entities.TblBill, error) {
	// transactional
	tx, err := db.Begin(ctx)
	if err != nil {
		return entities.TblBill{}, err
	}
	qtx := q.WithTx(tx)
	// insert tbl_bill
	var bill entities.TblBill

	var totalBillAmount int64
	for _, v := range payload.Services {
		var serviceDetail entities.TblServiceDetail
		result := int64(v.Quantity) * serviceDetail.UnitPrice
		row := qtx.db.QueryRow(ctx, createServices, serviceDetail.ServiceNo, v.Quantity, result)
		row.Scan(
			&v.ID,
			&v.ServiceNo,
			&v.Quantity,
			&v.Subtotal,
		)
		bill.Services = append(bill.Services, v)
		totalBillAmount += v.Subtotal
	}

	row := qtx.db.QueryRow(ctx, createBill, payload.DateIn, time.DateTime, payload.Customer.ID, bill.Services, totalBillAmount)
	err = row.Scan(
		&bill.ID,
		&bill.DateIn,
		&bill.DateFinished,
		&bill.Services,
		&bill.BillAmount,
	)
	if err != nil {
		return entities.TblBill{}, err
	}
	return bill, err
}
