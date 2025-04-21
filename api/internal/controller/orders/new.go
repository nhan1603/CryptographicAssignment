package orders

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderByID(ctx context.Context, id int) (model.Order, error)
	GetUserOrders(ctx context.Context, userID int) ([]model.Order, error)
	UpdateOrderStatus(ctx context.Context, id int, status string) error
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
