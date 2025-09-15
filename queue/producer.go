package queue

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	SyncProducer sarama.SyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{SyncProducer: prod}, nil
}

func (p *Producer) PublishCustomerDeactivated(id string) error {
	msg := &sarama.ProducerMessage{
		Topic: "customer-deactivated",
		Value: sarama.StringEncoder(fmt.Sprintf("Customer deactivated: ID=%s", id)),
	}
	_, _, err := p.SyncProducer.SendMessage(msg)
	return err
}
