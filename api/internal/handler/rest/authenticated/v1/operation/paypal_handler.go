package operation

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/iam"
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

type CreatePayPalOrderRequest struct {
	Items []model.OrderItem `json:"items,omitempty"`
}

type CreatePayPalOrderResponse struct {
	PayPalOrderID string `json:"paypal_order_id"`
	OrderID       int    `json:"order_id"`
}

func (h Handler) CreatePayPalOrder() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		var userData iam.HostProfile
		ctxUserValue := ctx.Value(iam.UserProfileKey)
		if ctxUserValue != nil {
			userData = ctxUserValue.(iam.HostProfile)
		} else {
			return webErrInternalServer
		}
		userID := userData.ID

		var req CreatePayPalOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return err
		}

		// validate the basic logic of the request itself
		if len(req.Items) <= 0 {
			return webErrInvalidRequest
		}

		for _, item := range req.Items {
			if item.ID < 0 || item.Quantity <= 0 {
				return webErrInvalidRequest
			}
		}

		paypalOrderId, orderID, err := h.orderCtrl.CreatePaypalOrder(ctx, model.Order{
			UserID:      userID,
			TotalAmount: float64(len(req.Items)),
			Items:       req.Items,
		})
		if err != nil {
			log.Printf("Error from controller: %+v\n", err)
			return webInternalSerror
		}

		httpserver.RespondJSON(w, CreatePayPalOrderResponse{
			PayPalOrderID: paypalOrderId,
			OrderID:       orderID,
		})

		return nil
	})
}

// CaptureResponse represents result of capturing order
type CaptureResponse struct {
	Success bool `json:"success"`
}

type CaptureRequest struct {
	PaypalID string `json:"paypal_order_id"`
	OrderID  int    `json:"order_id"`
}

func (h Handler) CapturePayPalOrder() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		var req CaptureRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return err
		}

		// validate request
		if req.OrderID <= 0 || req.PaypalID == "" {
			return webErrInvalidOrder
		}

		err := h.orderCtrl.CapturePaypalOrder(ctx, req.PaypalID, req.OrderID)
		if err != nil {
			log.Printf("Error from controller: %+v\n", err)
			return webInternalSerror
		}

		httpserver.RespondJSON(w, CaptureResponse{
			Success: true,
		})

		return nil
	})
}
