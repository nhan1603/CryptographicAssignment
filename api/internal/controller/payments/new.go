package payments

import (
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// CreatePayment(payment *model.PayPalTransaction) error
	// GetPaymentByOrderID(orderID int64) (*model.PayPalTransaction, error)
	// UpdatePaymentStatus(id int64, status string) error
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
