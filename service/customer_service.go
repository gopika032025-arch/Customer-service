package service

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gopika032025-arch/Customer-service/kafka"
	"github.com/gopika032025-arch/Customer-service/models"
)

type CustomerService struct {
	Producer sarama.SyncProducer
}

func (s *CustomerService) CreateCustomer(c models.Customer) error {
	// Mark customer as active
	c.Active = true

	// Create Kafka event
	event := kafka.CustomerEvent{
		ID:    strconv.Itoa(c.ID),
		Email: c.Email,
		Type:  "created",
	}

	// Publish event
	if err := kafka.PublishCustomerEvent(s.Producer, kafka.CustomerCreatedTopic, event); err != nil {
		log.Printf("Failed to publish event: %v", err)
		return err
	}

	return nil
}

func (s *CustomerService) DeactivateCustomer(c *models.Customer) error {
	c.Active = false

	event := kafka.CustomerEvent{
		ID:    strconv.Itoa(c.ID),
		Email: c.Email,
		Type:  "deactivated",
	}

	if err := kafka.PublishCustomerEvent(s.Producer, kafka.CustomerDeactivatedTopic, event); err != nil {
		log.Printf("Failed to publish event: %v", err)
		return err
	}

	return nil
}
