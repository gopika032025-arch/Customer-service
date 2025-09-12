package api

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
)

type Handler struct {
	db       *sql.DB
	producer sarama.SyncProducer
}

func (h *Handler) DeactivateHandler() func(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func (h *Handler) UpdateHandler() func(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func (h *Handler) CreateHandler() func(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func NewHandler(db *sql.DB, producer sarama.SyncProducer) *Handler {
	return &Handler{db: db, producer: producer}
}

func RegisterRoutes(db *sql.DB, producer sarama.SyncProducer) {
	h := NewHandler(db, producer)

	http.HandleFunc("/customer/create", h.CreateHandler())
	http.HandleFunc("/customer/update", h.UpdateHandler())
	http.HandleFunc("/customer/deactivate", h.DeactivateHandler())
}
