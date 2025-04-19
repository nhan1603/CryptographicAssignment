package menu

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	GetAll(context.Context) ([]model.MenuItem, error)
	GetByID(ctx context.Context, id int) (model.MenuItem, error)
	Create(ctx context.Context, item model.MenuItem) error
	Update(ctx context.Context, item model.MenuItem) error
	// Delete(ctx context.Context, id int64) error
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
