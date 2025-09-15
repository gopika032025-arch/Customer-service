package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/IBM/sarama"
)

func main() {
	dsn := "username:Abhi@123@tcp(localhost(127.0.0.1)customerdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	producer, err := initKafkaProducer()
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer producer.Close()
	// api.RegisterRoutes(db, producer)
	log.Println("Server starting on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func initKafkaProducer() (sarama.SyncProducer, error) {
	brokerList := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
