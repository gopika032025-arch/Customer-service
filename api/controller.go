package api

import (
	"encoding/json"
	"net/http"
)

type updateEmailReq struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Handler struct {
	biz IBizLogic
}

func NewHandler(biz IBizLogic) *Handler {
	return &Handler{biz: biz}
}

func (h Handler) UpdateEmailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Only PUT allowed", http.StatusMethodNotAllowed)
			return
		}

		var req updateEmailReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		err := h.biz.UpdateCustomerEmailLogic(req.ID, req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer email updated successfully"))
	}
}
