package api

import (
	"database/sql"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/juhithasabbineni0320/customer-service/dataservice"
	"github.com/juhithasabbineni0320/customer-service/models"
	"github.com/juhithasabbineni0320/customer-service/queue"
)

type IBizLogic interface {
	CreateCustomerLogic(customer models.CreateCustomerRequest) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, producer sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: producer}
}

func (bl *BizLogic) CreateCustomerLogic(customer models.CreateCustomerRequest) error {
	if customer.CustomerID == 0 {
		return fmt.Errorf("CustomerID is required")
	}
	if customer.Email == "" {
		return fmt.Errorf("email is required")
	}
	if customer.Password == "" {
		return fmt.Errorf("password is required")
	}

	if err := dataservice.CreateCustomer(bl.DB, customer); err != nil {
		return err
	}
	message := fmt.Sprintf("customerID: %d email: %s password: %s", customer.CustomerID, customer.Email, customer.Password)
	err := queue.ProduceKafkaMessage("customer_service", message, bl.Producer)
	if err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}

	return nil
}
