package models

type Customer struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}
