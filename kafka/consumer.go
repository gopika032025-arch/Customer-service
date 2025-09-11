
package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type customerEventHandler struct{}

func (h *customerEventHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *customerEventHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *customerEventHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		session.MarkMessage(msg, "")
	}
	return nil
}

func ConsumeCustomerEvents(consumerGroup sarama.ConsumerGroup, topics []string) {
	handler := &customerEventHandler{}
	for {
		err := consumerGroup.Consume(nil, topics, handler)
		if err != nil {
			log.Printf("Error consuming: %v\n", err)
		}
	}
}

