package models

type UpdateEmailReq struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
