package api

import (
	"database/sql"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gopika032025-arch/Customer-service/dataservice"
)

type IBizLogic interface {
	DeactivateCustomerLogic(id int) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, prod sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: prod}
}

func (bl *BizLogic) DeactivateCustomerLogic(id int) error {
	// Call the dataservice to deactivate in DB
	err := dataservice.DeactivateCustomer(bl.DB, id)
	if err != nil {
		return err
	}

	// Publish Kafka event
	msg := &sarama.ProducerMessage{
		Topic: "customer-events",
		Value: sarama.StringEncoder(fmt.Sprintf("Customer deactivated: %d", id)),
	}
	_, _, _ = bl.Producer.SendMessage(msg)

	return nil
}
