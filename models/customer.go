package models

type CreateCustomerRequest struct {
	CustomerID int `json:"customer_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}