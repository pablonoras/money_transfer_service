package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pablonoras/money_transfer_service/internal/core/ports"
	"net/http"
)

type HTTPHandler struct{
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (hdl *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "user_id")
	if len(userID) == 0 {
		http.Error(w, "Invalid user_id param,", 400)
	}

	user, err := hdl.service.GetBalance(ctx, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Get user error: %v", err.Error()), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*user)
}