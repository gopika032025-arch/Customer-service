package main

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/juhithasabbineni0320/customer-service/api"
)

func main() {
	db, _ := sql.Open("mysql", "user:pass@/dbname")
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	api.RegisterRoutes(db, producer)
	http.ListenAndServe(":8082", nil)
}
