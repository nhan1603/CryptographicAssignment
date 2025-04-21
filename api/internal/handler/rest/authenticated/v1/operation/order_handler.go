package operation

import (
	"encoding/json"
	"net/http"

	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

// CreateOrderResponse represents result of creating order
type CreateOrderResponse struct {
	Success bool        `json:"success"`
	Data    model.Order `json:"data"`
}

type CreateOrderRequest struct {
	UserID      int64             `json:"user_id"`
	TotalAmount float64           `json:"total_amount"`
	Items       []model.OrderItem `json:"items,omitempty"`
}

func (h Handler) CreateOrder() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {

		var req CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return err
		}

		if req.UserID < 0 || req.TotalAmount <= 0 {
			return webErrInvalidRequest
		}

		for _, item := range req.Items {
			if item.ID < 0 || item.Quantity <= 0 || item.UnitPrice <= 0 {
				return webErrInvalidRequest
			}
		}

		orderData, err := h.orderCtrl.CreateOrder(r.Context(), model.Order{
			UserID:      req.UserID,
			TotalAmount: req.TotalAmount,
			Items:       req.Items,
		})
		if err != nil {
			return webInternalSerror
		}

		httpserver.RespondJSON(w, CreateOrderResponse{
			Success: true,
			Data:    orderData,
		})

		return nil
	})
}
