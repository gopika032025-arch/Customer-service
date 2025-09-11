package main

import (
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/gopika032025-arch/Customer-service/models"
	"github.com/gopika032025-arch/Customer-service/service"
)

func main() {
	// --- 1. Kafka producer setup ---
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer producer.Close()

	// --- 2. Initialize CustomerService ---
	custService := service.CustomerService{Producer: producer}

	// --- 3. Create a customer ---
	customer := models.Customer{
		ID:       1,
		Email:    "test@example.com",
		Password: "password123",
	}

	if err := custService.CreateCustomer(customer); err != nil {
		log.Printf("Error creating customer: %v", err)
	}

	// --- 4. Deactivate customer after some time ---
	time.Sleep(2 * time.Second)

	if err := custService.DeactivateCustomer(&customer); err != nil {
		log.Printf("Error deactivating customer: %v", err)
	}

	// --- 5. Optional: Consume events ---
	// consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "customer-group", nil)
	// if err != nil {
	// 	log.Fatalf("Failed to create consumer group: %v", err)
	// }
	// defer consumerGroup.Close()
	// go kafka.ConsumeCustomerEvents(consumerGroup, []string{kafka.CustomerCreatedTopic, kafka.CustomerDeactivatedTopic})

	log.Println("Customer events sent successfully!")
}
