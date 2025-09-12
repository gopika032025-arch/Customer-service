package models

type CreateCustomerRequest struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UpdateCustomerRequest struct {
	CustomerID string `json:"customer_id"`
	NewEmail   string `json:"new_email"`
}

type DeactivateCustomerRequest struct {
	CustomerID string `json:"customer_id"`
}
