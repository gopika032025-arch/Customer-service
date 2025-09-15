package dataservice

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("customer not found")

func DeactivateCustomer(db *sql.DB, id string) error {
	query := "UPDATE customers SET active = FALSE WHERE id = ?"
	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
