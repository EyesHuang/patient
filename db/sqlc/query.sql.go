// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const getAll = `-- name: GetAll :many
SELECT
    patient.id AS patient_id,
    patient.name AS patient_name,
    medical_order.message AS order_message
FROM
    patient
        JOIN medical_order ON patient.order_id = medical_order.id
`

type GetAllRow struct {
	PatientID    int32
	PatientName  sql.NullString
	OrderMessage sql.NullString
}

func (q *Queries) GetAll(ctx context.Context) ([]GetAllRow, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllRow
	for rows.Next() {
		var i GetAllRow
		if err := rows.Scan(&i.PatientID, &i.PatientName, &i.OrderMessage); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE medical_order
SET message = $1
    WHERE id = $2
`

type UpdateOrderParams struct {
	Message sql.NullString
	ID      int32
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder, arg.Message, arg.ID)
	return err
}