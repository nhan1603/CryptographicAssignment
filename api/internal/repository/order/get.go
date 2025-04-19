package order

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r impl) GetByID(ctx context.Context, id int) (model.Order, error) {
	data, err := dbmodel.Orders(dbmodel.OrderWhere.ID.EQ(id)).One(ctx, r.dbConn)
	if err != nil {
		return model.Order{}, err
	}

	// Get order items for this order
	items, err := dbmodel.OrderItems(dbmodel.OrderItemWhere.OrderID.EQ((id)),
		qm.Load(dbmodel.OrderItemRels.MenuItem)).All(ctx, r.dbConn)
	if err != nil {
		return model.Order{}, err
	}

	// Convert order items to model
	orderItems := []model.OrderItem{}
	for _, item := range items {
		product := item.R.MenuItem
		orderItems = append(orderItems, model.OrderItem{
			ID:         int64(item.ID),
			OrderID:    int64(item.OrderID),
			MenuItemID: int64(item.MenuItemID),
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			Subtotal:   item.Subtotal,
			MenuItem: model.MenuItem{
				ID:          int64(data.ID),
				Name:        product.Name,
				Description: product.Description.String,
				Category:    product.Category.String,
				Price:       product.Price,
				ImageUrl:    product.ImageURL.String,
				IsAvailable: product.IsAvailable.Bool,
			},
		})
	}

	return model.Order{
		ID:          int64(data.ID),
		UserID:      int64(data.UserID),
		TotalAmount: data.TotalAmount,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Time,
		Items:       orderItems,
	}, nil
}

func (r impl) GetByUserID(ctx context.Context, userID int) ([]model.Order, error) {
	datas, err := dbmodel.Orders(dbmodel.OrderWhere.UserID.EQ(userID)).All(ctx, r.dbConn)
	if err != nil {
		return nil, err
	}

	result := []model.Order{}

	for _, data := range datas {
		// Get order items for each order
		items, err := dbmodel.OrderItems(dbmodel.OrderItemWhere.OrderID.EQ(data.ID)).All(ctx, r.dbConn)
		if err != nil {
			return nil, err
		}

		// Convert order items to model
		orderItems := []model.OrderItem{}
		for _, item := range items {
			orderItems = append(orderItems, model.OrderItem{
				ID:         int64(item.ID),
				OrderID:    int64(item.OrderID),
				MenuItemID: int64(item.MenuItemID),
				Quantity:   item.Quantity,
				UnitPrice:  item.UnitPrice,
				Subtotal:   item.Subtotal,
			})
		}

		result = append(result, model.Order{
			ID:          int64(data.ID),
			UserID:      int64(data.UserID),
			TotalAmount: data.TotalAmount,
			Status:      data.Status,
			CreatedAt:   data.CreatedAt.Time,
			Items:       orderItems,
		})
	}

	return result, nil
}
