package menus

import (
	"errors"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

func (c impl) GetAllItems() ([]model.MenuItem, error) {
	return c.repo.GetAll()
}

func (c impl) GetItemByID(id int64) (model.MenuItem, error) {
	item, err := c.repo.GetByID(id)
	if err != nil {
		return model.MenuItem{}, err
	}
	if item == nil {
		return model.MenuItem{}, errors.New("menu item not found")
	}
	return item, nil
}

func (c impl) CreateItem(item model.MenuItem) error {
	if item.Name == "" || item.Price <= 0 {
		return errors.New("invalid menu item data")
	}
	return c.repo.Create(item model.MenuItem)
}

func (c impl) UpdateItem(item model.MenuItem) error {
	if item.ID <= 0 || item.Name == "" || item.Price <= 0 {
		return errors.New("invalid menu item data")
	}
	return c.repo.Update(item)
}

func (c impl) DeleteItem(id int64) error {
	return c.repo.Delete(id)
}
