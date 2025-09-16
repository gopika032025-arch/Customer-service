package api

import (
	"encoding/json"
	"net/http"

	"github.com/gopika032025-arch/Customer-service/models"
)

func (h Handler) DeactivateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.Customer
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		err := h.biz.DeactivateCustomerLogic(req.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer deactivated successfully"))
	}
}
