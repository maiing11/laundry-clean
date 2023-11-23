package repository

import (
	"context"
	"database/sql"
	"time"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/model/entities"
)

const createBill = `INSERT INTO tbl_bills (
	date_in, date_finished, customer_id, services, bill_amount
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, date_in, date_finished, customer_id, services, bill_amount
)`

const createServices = `INSERT INTO tbl_services service_no, quantity, subtotal
VALUES $1, $2, $3
RETURNING id, service_no, quantity, subtotal
`

func (q *Queries) CreateBill(ctx context.Context, db *sql.DB, payload entities.TblBill) (entities.TblBill, error) {
	// transactional
	tx, err := db.Begin()
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
		row := qtx.db.QueryRowContext(ctx, createServices, serviceDetail.ServiceNo, v.Quantity, result)
		row.Scan(
			&v.ID,
			&v.ServiceNo,
			&v.Quantity,
			&v.Subtotal,
		)
		bill.Services = append(bill.Services, v)
		totalBillAmount += v.Subtotal
	}

	row := qtx.db.QueryRowContext(ctx, createBill, payload.DateIn, time.DateTime, payload.Customer.ID, bill.Services, totalBillAmount)
	err = row.Scan(
		&bill.ID,
		&bill.DateIn,
		time.Now(),
		&bill.Services,
		&bill.BillAmount,
	)
	if err != nil {
		return entities.TblBill{}, err
	}
	return bill, err
}
