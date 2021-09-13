package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pablonoras/money_transfer_service/internal/core/ports"
	"net/http"
	"strconv"
)

type HTTPHandler struct{
	service ports.TransactionService
}

func NewTransactionHandler(service ports.TransactionService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (hdl *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "user_id")
	if len(userID) == 0 {
		http.Error(w, "Invalid user_id param,", 400)
	}

	transactions, err := hdl.service.Get(ctx, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Get transactions error: %v", err.Error()), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func (hdl *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "user_id")
	if len(userID) == 0 {
		http.Error(w, "Invalid user_id param,", 400)
	}

	receptorID := chi.URLParam(r,"receptor_id")

	if len(receptorID) == 0 {
		http.Error(w, "Url param 'receptor_id' not found ", 400)
	}

	amount := chi.URLParam(r,"amount")
	if len(amount) == 0 {
		http.Error(w, "Url param 'amount' not found ", 400)
	}

	amountFormatted , err := strconv.Atoi(amount)

	transaction, err := hdl.service.Create(ctx, userID, receptorID, amountFormatted)
	if err != nil {
		http.Error(w, fmt.Sprintf("Create transactions error: %v", err.Error()), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&transaction)
}