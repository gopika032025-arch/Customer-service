package api

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
)

type Handler struct {
	biz IBizLogic
}

func SetupRoutes(db *sql.DB, producer sarama.SyncProducer) {
	biz := NewBizLogic(db, producer)
	handler := Handler{biz: biz}

	http.HandleFunc("/customers/deactivate", handler.DeactivateHandler())
}
