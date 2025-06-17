package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ricardo-ronchini/budget-flow-app-go/auth"
)

type Expense struct {
	Name   string  `json:"name"`
	Amount float64 `json:"value"`
}

func Expenses(w http.ResponseWriter, r *http.Request) {
	// Recupera userID do context
	userID, ok := auth.GetUserID(r.Context())
	if !ok {
		http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	expenses := []Expense{
		{"Mercado", 100},
		{"Uber", 35},
	}

	json.NewEncoder(w).Encode(map[string]any{
		"user":     userID,
		"expenses": expenses,
	})
}
