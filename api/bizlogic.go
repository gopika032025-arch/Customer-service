package api

import (
	"github.com/gopika032025-arch/Customer-service/dataservice"

	"database/sql"

	"github.com/IBM/sarama"
)

type IBizLogic interface {
	DeactivateCustomerLogic(id string) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, prod sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: prod}
}

func (bl *BizLogic) DeactivateCustomerLogic(id string) error {
	err := dataservice.DeactivateCustomer(bl.DB, id)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "customer-events",
		Value: sarama.StringEncoder("Customer deactivated: " + id),
	}
	_, _, _ = bl.Producer.SendMessage(msg)

	return nil
}
