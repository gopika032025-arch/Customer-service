package dataservice

import (
	"database/sql"

	"github.com/juhithasabbineni0320/customer-service/models"
)

func CreateCustomer(db *sql.DB, customer models.CreateCustomerRequest) error {
	query := `
		INSERT INTO customers (customer_id, email, password)
		VALUES (?, ?, ?)`
	_, err := db.Exec(query, customer.CustomerID, customer.Email, customer.Password)
	return err
}
