package menu

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r impl) Create(ctx context.Context, item model.MenuItem) error {
	dbMenu := dbmodel.MenuItem{
		Name:        item.Name,
		Description: null.StringFrom(item.Description),
		Price:       item.Price,
		Category:    null.StringFrom(item.Category),
		IsAvailable: null.BoolFrom(true),
	}
	err := dbMenu.Insert(ctx, r.dbConn, boil.Infer())

	if err != nil {
		pkgerrors.WithStack(err)
	}
	return nil
}

func (r impl) Update(ctx context.Context, item model.MenuItem) error {
	dbMenu := dbmodel.MenuItem{
		Name:        item.Name,
		Description: null.StringFrom(item.Description),
		Price:       item.Price,
		ImageURL:    null.StringFrom(item.ImageUrl),
		Category:    null.StringFrom(item.Category),
		IsAvailable: null.BoolFrom(true),
	}
	err := dbMenu.Insert(ctx, r.dbConn, boil.Infer())

	if err != nil {
		pkgerrors.WithStack(err)
	}
	return nil
}
