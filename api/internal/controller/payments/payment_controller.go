package payments

import (
	"errors"
	"time"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

func (c impl) CreatePayment(payment *model.PayPalTransaction) error {
	if payment.OrderID <= 0 || payment.PaymentAmount <= 0 {
		return errors.New("invalid payment data")
	}

	// Verify order exists and amount matches
	order, err := c.orderRepo.GetByID(payment.OrderID)
	if err != nil {
		return err
	}
	if order == nil {
		return errors.New("order not found")
	}

	if order.TotalAmount != payment.PaymentAmount {
		return errors.New("payment amount does not match order total")
	}

	payment.PaymentStatus = model.PaymentStatusPending
	payment.PaymentDate = time.Now()

	if err := c.paymentRepo.Create(payment); err != nil {
		return err
	}

	// Update order status to paid
	order.Status = model.OrderStatusPaid
	return c.orderRepo.Update(order)
}

func (c impl) GetPaymentByOrderID(orderID int64) (*model.PayPalTransaction, error) {
	payment, err := c.paymentRepo.GetByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	if payment == nil {
		return nil, errors.New("payment not found")
	}
	return payment, nil
}

func (c impl) UpdatePaymentStatus(id int64, status string) error {
	if !isValidPaymentStatus(status) {
		return errors.New("invalid payment status")
	}

	payment, err := c.paymentRepo.GetByID(id)
	if err != nil {
		return err
	}
	if payment == nil {
		return errors.New("payment not found")
	}

	payment.PaymentStatus = status
	return c.paymentRepo.Update(payment)
}

func isValidPaymentStatus(status string) bool {
	validStatuses := []string{
		model.PaymentStatusPending,
		model.PaymentStatusCompleted,
		model.PaymentStatusFailed,
		model.PaymentStatusRefunded,
	}

	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}
