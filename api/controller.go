package api

import (
	"encoding/json"
	"net/http"

	"github.com/juhithasabbineni0320/customer-service/models"
)

type Handler struct {
	BizLogic *BizLogic
}

func (h *Handler) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var customer models.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.BizLogic.CreateCustomerLogic(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer created successfully"})
}
