package repository

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

const createCustomers = `-- name: CreateCustomers :one
INSERT INTO tbl_customers (
  name, address, phone_number
) VALUES (
  $1, $2, $3
)
RETURNING id, name, address, phone_number
`

func (q *Queries) CreateCustomers(ctx context.Context, arg entities.TblCustomer) (entities.TblCustomer, error) {
	row := q.db.QueryRowContext(ctx, createCustomers, arg.Name, arg.Address, arg.PhoneNumber)
	var i entities.TblCustomer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM tbl_customers
WHERE id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, name, address, phone_number FROM tbl_customers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, id string) (entities.TblCustomer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i entities.TblCustomer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
	)
	return i, err
}

const listCustomers = `-- name: ListCustomers :many
SELECT id, name, address, phone_number FROM tbl_customers
ORDER BY name
`

func (q *Queries) ListCustomers(ctx context.Context) ([]entities.TblCustomer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []entities.TblCustomer
	for rows.Next() {
		var i entities.TblCustomer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.PhoneNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCustomer = `-- name: UpdateCustomer :exec
UPDATE tbl_customers
  set name = $2,
  address = $3,
  phone_number = $4
WHERE id = $1
`

func (q *Queries) UpdateCustomer(ctx context.Context, arg entities.TblCustomer) error {
	_, err := q.db.ExecContext(ctx, updateCustomer,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.PhoneNumber,
	)
	return err
}
