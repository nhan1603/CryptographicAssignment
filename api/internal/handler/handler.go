package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	menuService    MenuService
	orderService   OrderService
	paymentService PaymentService
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewHandler(menuService MenuService, orderService OrderService, paymentService PaymentService) *Handler {
	return &Handler{
		menuService:    menuService,
		orderService:   orderService,
		paymentService: paymentService,
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response := Response{
		Success: code >= 200 && code < 300,
		Data:    payload,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) respondWithError(w http.ResponseWriter, code int, message string) {
	response := Response{
		Success: false,
		Error:   message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
