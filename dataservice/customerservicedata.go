package dataservice

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("customer not found")

func UpdateCustomerEmail(db *sql.DB, id int, newEmail string) error {
	res, err := db.Exec("UPDATE customers SET email = ?, active = TRUE WHERE id = ?", newEmail, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}
