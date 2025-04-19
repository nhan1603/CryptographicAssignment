package menus

import (
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	GetAllItems() ([]model.MenuItem, error)
	GetItemByID(id int64) (model.MenuItem, error)
	CreateItem(item model.MenuItem) error
	UpdateItem(item model.MenuItem) error
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
