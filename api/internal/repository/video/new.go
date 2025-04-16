package video

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	ShareVideo(ctx context.Context, sharingInput model.VideoShare) error
	RetrieveSharedVideo(ctx context.Context) (model.ListSharedVideo, error)
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
