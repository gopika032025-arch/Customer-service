package api

import (
	"database/sql"
	"fmt"

	"github.com/juhithasabbineni0320/customer-service/dataservice"
	"github.com/juhithasabbineni0320/customer-service/models"
)

type IBizLogic interface {
	CreateCustomerLogic(customer models.CreateCustomerRequest) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) CreateCustomerLogic(customer models.CreateCustomerRequest) error {
	if customer.CustomerID == "" {
		return fmt.Errorf("CustomerID is required")
	}
	if customer.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if customer.Password == "" {
		return fmt.Errorf("Password is required")
	}

	if err := dataservice.CreateCustomer(bl.DB, customer); err != nil {
		return err
	}

	return nil
}
