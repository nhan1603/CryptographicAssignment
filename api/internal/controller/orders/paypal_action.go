package orders

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	crypto_helper "github.com/nhan1603/CryptographicAssignment/api/internal/pkg/cryptos"
	"github.com/plutov/paypal/v4"
)

func (c impl) CreatePaypalOrder(ctx context.Context, order model.Order) (string, int, error) {
	if order.UserID <= 0 || len(order.Items) == 0 {
		return "", 0, errors.New("invalid order data")
	}

	// Calculate total amount and validate items
	var totalAmount float64
	for indx, item := range order.Items {
		menuItem, err := c.repo.Menu().GetByID(ctx, int(item.MenuItemID))
		if err != nil {
			return "", 0, err
		}

		if !menuItem.IsAvailable {
			return "", 0, errors.New("menu item not available")
		}

		item.UnitPrice = menuItem.Price
		item.Subtotal = menuItem.Price * float64(item.Quantity)
		totalAmount += item.Subtotal
		order.Items[indx] = item
	}

	// Create PayPal client
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")
	client, err := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)
	if err != nil {
		log.Printf("Err initialize paypal client: %+v\n", err)
		return "", 0, err
	}
	_, err = client.GetAccessToken(context.Background())
	if err != nil {
		log.Printf("PayPal auth error: %+v\n", err)
		return "", 0, err
	}

	// Create PayPal order
	paypalOrder, err := client.CreateOrder(context.Background(), paypal.OrderIntentCapture,
		[]paypal.PurchaseUnitRequest{
			{
				Amount: &paypal.PurchaseUnitAmount{
					Currency: "GBP",
					Value:    fmt.Sprintf("%.2f", totalAmount),
				},
				Description: "Food Order",
			},
		}, nil,
		&paypal.ApplicationContext{
			BrandName:  "University Catering System",
			UserAction: "PAY_NOW",
			ReturnURL:  "https://example.com/return",
			CancelURL:  "https://example.com/cancel",
		})
	if err != nil {
		return "", 0, err
	}

	// Save order in DB with status "pending"
	order.TotalAmount = totalAmount
	order.Status = model.OrderStatusPending
	order.CreatedAt = time.Now()

	var orderID int
	if order.ID == 0 {
		orderID, err = c.repo.Order().Create(ctx, order)
	} else {
		orderID, err = c.repo.Order().Update(ctx, order)
	}
	if err != nil {
		log.Printf("Err creating internal order: %+v\n", err)
		return "", 0, err
	}

	return paypalOrder.ID, orderID, nil
}

func (c impl) CapturePaypalOrder(ctx context.Context, payPalID string, id int) error {
	order, err := c.repo.Order().GetByID(ctx, id)
	if err != nil {
		// if no internal order exist, we abort
		return err
	}
	// Create PayPal client
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")
	client, err := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)
	if err != nil {
		log.Printf("Err initialize paypal client: %+v\n", err)
		return err
	}

	// Call PayPal API to capture the order
	captureResult, err := client.CaptureOrder(ctx, payPalID, paypal.CaptureOrderRequest{})
	if err != nil {
		log.Printf("Err capturing order status: %+v\n", err)
		return err
	}

	// Simultaneously capture the order information in a separate thread
	go func() {
		var PaymentAmount string
		if len(captureResult.PurchaseUnits) > 0 && len(captureResult.PurchaseUnits[0].Payments.Captures) > 0 {
			PaymentAmount = captureResult.PurchaseUnits[0].Payments.Captures[0].Amount.Value
		}

		// Encrypt the payer's email to enhance security
		encryptionKey := os.Getenv("CYPHER_KEY")
		encryptedEmail, err := crypto_helper.EncryptMessage([]byte(encryptionKey), captureResult.Payer.EmailAddress)
		if err != nil {
			log.Printf("Err encrypting email for order with ID %d and Paypal order ID %s: %+v\n", order.ID, payPalID, err)
			return
		}

		err = c.repo.Payment().Create(context.Background(), model.PayPalTransaction{
			OrderID:             order.ID,
			PayPalTransactionID: payPalID,
			PaymentStatus:       captureResult.Status,
			Currency:            "GBP",
			PaymentAmount:       PaymentAmount,
			PayerEmail:          encryptedEmail,
		})
		if err != nil {
			log.Printf("Err saving order with ID %d and Paypal order ID %s capture result: %+v\n", order.ID, payPalID, err)
		}
	}()

	// Verify the status of the order
	if captureResult.Status != "COMPLETED" {
		log.Printf("Error payment not completed")
		return errors.New("error payment not completed")
	}

	//Verify the captured amount matches your order
	if len(captureResult.PurchaseUnits) > 0 &&
		len(captureResult.PurchaseUnits[0].Payments.Captures) > 0 {
		capturedAmount := captureResult.PurchaseUnits[0].Payments.Captures[0].Amount.Value
		if capturedAmount != fmt.Sprintf("%.2f", order.TotalAmount) {
			log.Printf("Captured amount mismatch: PayPal=%s, Internal=%.2f", capturedAmount, order.TotalAmount)
			return errors.New("captured amount does not match order total")
		}
	}

	// Update order status to paid
	return c.repo.Order().UpdateStatus(ctx, id, model.OrderStatusPaid)
}
