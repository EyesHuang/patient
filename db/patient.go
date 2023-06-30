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
