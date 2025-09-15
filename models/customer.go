package models

type Customer struct {
	ID       int  `json:"id"`
	Active   bool `json:"active"`
	Inactive bool `json:"inactive"`
}
