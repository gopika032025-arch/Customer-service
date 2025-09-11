package kafka

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func PublishCustomerEvent(producer sarama.SyncProducer, topic string, event CustomerEvent) error {
	messageBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Published event to topic %s: %+v\n", topic, event)
	return nil
}
