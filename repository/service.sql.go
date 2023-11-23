package repository

import (
	"context"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/model/entities"
)

const createServiceDetails = `-- name: CreateServiceDetails :one
INSERT INTO tbl_service_details (
  service_name, units, unit_price
) VALUES (
  $1, $2, $3
)
RETURNING service_no, service_name, units, unit_price
`

func (q *Queries) CreateServiceDetails(ctx context.Context, arg entities.TblServiceDetail) (entities.TblServiceDetail, error) {
	row := q.db.QueryRowContext(ctx, createServiceDetails, arg.ServiceName, arg.Units, arg.UnitPrice)
	var i entities.TblServiceDetail
	err := row.Scan(
		&i.ServiceNo,
		&i.ServiceName,
		&i.Units,
		&i.UnitPrice,
	)
	return i, err
}

const deleteServiceDetail = `-- name: DeleteServiceDetail :exec
DELETE FROM tbl_service_details
WHERE service_no = $1
`

func (q *Queries) DeleteServiceDetail(ctx context.Context, serviceNo int) error {
	_, err := q.db.ExecContext(ctx, deleteServiceDetail, serviceNo)
	return err
}

const getServiceDetails = `-- name: GetServiceDetails :one
SELECT service_no, service_name, units, unit_price FROM tbl_service_details
WHERE service_no = $1 LIMIT 1
`

func (q *Queries) GetServiceDetails(ctx context.Context, serviceNo int) (entities.TblServiceDetail, error) {
	row := q.db.QueryRowContext(ctx, getServiceDetails, serviceNo)
	var i entities.TblServiceDetail
	err := row.Scan(
		&i.ServiceNo,
		&i.ServiceName,
		&i.Units,
		&i.UnitPrice,
	)
	return i, err
}

const listServiceDetails = `-- name: ListServiceDetails :many
SELECT service_no, service_name, units, unit_price FROM tbl_service_details
ORDER BY service_name
`

func (q *Queries) ListServiceDetails(ctx context.Context) ([]entities.TblServiceDetail, error) {
	rows, err := q.db.QueryContext(ctx, listServiceDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []entities.TblServiceDetail
	for rows.Next() {
		var i entities.TblServiceDetail
		if err := rows.Scan(
			&i.ServiceNo,
			&i.ServiceName,
			&i.Units,
			&i.UnitPrice,
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

const updateServiceDetails = `-- name: UpdateServiceDetails :exec
UPDATE tbl_service_details
  set service_name = $2,
  units = $3,
  unit_price = $4
WHERE service_no = $1
`

func (q *Queries) UpdateServiceDetails(ctx context.Context, arg entities.TblServiceDetail) error {
	_, err := q.db.ExecContext(ctx, updateServiceDetails,
		arg.ServiceNo,
		arg.ServiceName,
		arg.Units,
		arg.UnitPrice,
	)
	return err
}
