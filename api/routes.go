package api

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
)

func RegisterRoutes(db *sql.DB, producer sarama.SyncProducer) {
	 h := &Handler{}
	http.HandleFunc("/create", h.CreateCustomerHandler())
}

func NewHandler(db *sql.DB, producer sarama.SyncProducer) any {
	panic("unimplemented")
}
