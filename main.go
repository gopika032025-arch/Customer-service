package main

import (
	"github.com/gopika032025-arch/Customer-service/api"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := ""
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	fmt.Println(" MySQL connected")

	producer, err := initKafkaProducer()
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer producer.Close()
	fmt.Println(" Kafka producer ready")

	api.SetupRoutes(db, producer)

	log.Println(" Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initKafkaProducer() (sarama.SyncProducer, error) {
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokers, config)
}
