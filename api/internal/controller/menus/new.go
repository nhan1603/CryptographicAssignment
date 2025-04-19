package menus

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	GetAllItems(ctx context.Context) ([]model.MenuItem, error)
	GetItemByID(ctx context.Context, id int) (model.MenuItem, error)
	CreateItem(ctx context.Context, item model.MenuItem) error
	UpdateItem(ctx context.Context, item model.MenuItem) error
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
