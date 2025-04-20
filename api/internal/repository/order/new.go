package order

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	Create(ctx context.Context, order model.Order) (int, error)
	Update(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, orderId int, status string) error
	GetByID(ctx context.Context, id int) (model.Order, error)
	GetByUserID(ctx context.Context, userID int) ([]model.Order, error)
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
