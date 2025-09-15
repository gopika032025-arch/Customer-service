package api

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
)

func RegisterRoutes(db *sql.DB, producer sarama.SyncProducer) {
	biz := NewBizLogic(db)
	handler := NewHandler(biz)

	http.HandleFunc("/update-email", handler.UpdateEmailHandler())
}
