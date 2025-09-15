package api

import (
	"customer-service/dataservice"
	"database/sql"
)

type IBizLogic interface {
	UpdateCustomerEmailLogic(id int, newEmail string) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (b *BizLogic) UpdateCustomerEmailLogic(id int, newEmail string) error {
	return dataservice.UpdateCustomerEmail(b.DB, id, newEmail)
}
