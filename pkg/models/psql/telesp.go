package psql

import (
	"database/sql"
	"telesp/pkg/models"
)

type TeleSp struct {
	DB *sql.DB
}

// Insert
func (m *TeleSp) Insert() error {
	return nil
}

// Find - select func for looking for based offered params
func (m *TeleSp) Get() (*models.PersonData, error) {
	return nil, nil
}

//