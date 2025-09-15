package models

type Customer struct {
	ID       int  `json:"id"`
	Active   bool `json:"active"`
	Inactive bool `json:"inactive"`
}

type CreateCustomerRequest struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
type UpdateEmailReq struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
