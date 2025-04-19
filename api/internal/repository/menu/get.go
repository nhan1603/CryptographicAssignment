package menu

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/dbmodel"
)

func (r impl) GetAll(ctx context.Context) ([]model.MenuItem, error) {
	datas, err := dbmodel.MenuItems().All(ctx, r.dbConn)
	if err != nil {
		return nil, err
	}

	result := []model.MenuItem{}

	for _, data := range datas {
		result = append(result, model.MenuItem{
			ID:          int64(data.ID),
			Name:        data.Name,
			Description: data.Description.String,
			Category:    data.Category.String,
			ImageUrl:    data.ImageURL.String,
			Price:       data.Price,
			IsAvailable: data.IsAvailable.Bool,
			CreatedAt:   data.CreatedAt.Time,
		})
	}

	return result, nil
}

func (r impl) GetByID(ctx context.Context, id int) (model.MenuItem, error) {
	data, err := dbmodel.MenuItems(dbmodel.MenuItemWhere.ID.EQ(id)).One(ctx, r.dbConn)
	if err != nil {
		return model.MenuItem{}, err
	}

	return model.MenuItem{
		ID:          int64(data.ID),
		Name:        data.Name,
		Description: data.Description.String,
		Category:    data.Category.String,
		Price:       data.Price,
		ImageUrl:    data.ImageURL.String,
		IsAvailable: data.IsAvailable.Bool,
		CreatedAt:   data.CreatedAt.Time,
	}, nil
}
