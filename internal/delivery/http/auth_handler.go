package http

import (
	"encoding/json"
	"net/http"

	"github.com/waldy-ctt/core-link-be/internal/domain/usecase"
)

type AuthHandler struct {
	signupUC usecase.SignupUseCase
}

func NewAuthHandler(signupUC usecase.SignupUseCase) *AuthHandler {
	return &AuthHandler{
		signupUC: signupUC,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var input usecase.SignupInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON Format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := h.signupUC.Execute(ctx, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User registered successfully"}`))
}
