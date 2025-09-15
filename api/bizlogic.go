package api

import (
	"customer-service/dataservice"
	"database/sql"
)

type IBizLogic interface {
	UpdateCustomerEmailLogic(id string, newEmail string) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) UpdateCustomerEmailLogic(id string, newEmail string) error {
	return dataservice.UpdateCustomerEmail(bl.DB, id, newEmail)
}
