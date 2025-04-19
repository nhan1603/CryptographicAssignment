package menus

import (
	"context"
	"errors"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

func (c impl) GetAllItems(ctx context.Context) ([]model.MenuItem, error) {
	return c.repo.Menu().GetAll(ctx)
}

func (c impl) GetItemByID(ctx context.Context, id int) (model.MenuItem, error) {
	item, err := c.repo.Menu().GetByID(ctx, id)
	if err != nil {
		return model.MenuItem{}, err
	}
	return item, nil
}

func (c impl) CreateItem(ctx context.Context, item model.MenuItem) error {
	if item.Name == "" || item.Price <= 0 {
		return errors.New("invalid menu item data")
	}
	return c.repo.Menu().Create(ctx, item)
}

func (c impl) UpdateItem(ctx context.Context, item model.MenuItem) error {
	if item.ID <= 0 || item.Name == "" || item.Price <= 0 {
		return errors.New("invalid menu item data")
	}
	return c.repo.Menu().Update(ctx, item)
}
