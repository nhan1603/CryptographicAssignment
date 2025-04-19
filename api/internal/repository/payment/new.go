package payment

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// Create(payment model.PayPalTransaction) error
	// GetByID(id int64) (model.PayPalTransaction, error)
	// GetByOrderID(orderID int64) (model.PayPalTransaction, error)
	// Update(payment model.PayPalTransaction) error
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}

}

type impl struct {
	dbConn boil.ContextExecutor
}
