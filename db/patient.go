package db

import (
	"context"
	"database/sql"

	patient "manage-patinets"
	"manage-patinets/db/sqlc"
)

type PatientCommands interface {
	WithTx(tx *sql.Tx) *sqlc.Queries
	GetAll(ctx context.Context) ([]sqlc.GetAllRow, error)
	UpdateOrder(ctx context.Context, arg sqlc.UpdateOrderParams) error
}

type PatientRepository struct {
	db       *sql.DB
	commands PatientCommands
}

func NewPatientRepositoryWithCommand(db *sql.DB, commands PatientCommands) *PatientRepository {
	return &PatientRepository{
		db:       db,
		commands: commands,
	}
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return NewPatientRepositoryWithCommand(db, sqlc.New(db))
}

func (pr *PatientRepository) GetAllPatients(ctx context.Context) (*[]patient.Patient, error) {
	pRows, err := pr.commands.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	patients := make([]patient.Patient, 0)

	for _, pRow := range pRows {
		order := patient.Order{
			ID:      int(pRow.OrderID),
			Message: pRow.OrderMessage.String,
		}

		p := patient.Patient{
			ID:    int(pRow.PatientID),
			Name:  pRow.PatientName.String,
			Order: order,
		}
		patients = append(patients, p)
	}

	return &patients, nil
}

func (pr *PatientRepository) Update(ctx context.Context, order *patient.Order) (*patient.Order, error) {
	params := sqlc.UpdateOrderParams{
		Message: sql.NullString{
			String: order.Message,
			Valid:  true,
		},
		ID: int32(order.ID),
	}

	if err := pr.commands.UpdateOrder(ctx, params); err != nil {
		return nil, err
	}

	return order, nil
}
