package order

import (
	"context"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/dbmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r impl) Create(ctx context.Context, order model.Order) (int, error) {
	orderDb := dbmodel.Order{
		UserID:      int(order.UserID),
		TotalAmount: order.TotalAmount,
	}

	err := orderDb.Insert(ctx, r.dbConn, boil.Infer())
	if err != nil {
		return 0, err
	}

	orderId := orderDb.ID
	for _, item := range order.Items {
		itemDb := dbmodel.OrderItem{
			OrderID:    orderId,
			MenuItemID: int(item.MenuItemID),
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			Subtotal:   float64(item.Quantity) * item.UnitPrice,
		}

		err := itemDb.Insert(ctx, r.dbConn, boil.Infer())
		if err != nil {
			return 0, err
		}
	}
	return orderId, nil
}

func (r impl) Update(ctx context.Context, order model.Order) error {
	orderDb, err := dbmodel.Orders(dbmodel.OrderWhere.ID.EQ(int(order.ID)),
		qm.Load(dbmodel.OrderRels.OrderItems)).One(ctx, r.dbConn)
	if err != nil {
		return err
	}

	if orderDb.R != nil && orderDb.R.OrderItems != nil {
		for _, item := range orderDb.R.OrderItems {
			_, err = item.Delete(ctx, r.dbConn)
			if err != nil {
				return err
			}
		}
	}

	for _, item := range order.Items {
		itemDb := dbmodel.OrderItem{
			OrderID:    int(order.ID),
			MenuItemID: int(item.MenuItemID),
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			Subtotal:   float64(item.Quantity) * item.UnitPrice,
		}

		err := itemDb.Insert(ctx, r.dbConn, boil.Infer())
		if err != nil {
			return err
		}
	}

	orderDb.Status = order.Status
	orderDb.TotalAmount = order.TotalAmount

	cols := []string{
		dbmodel.OrderColumns.Status,
		dbmodel.OrderColumns.TotalAmount,
	}

	_, err = orderDb.Update(ctx, r.dbConn, boil.Whitelist(cols...))
	return err
}

func (r impl) UpdateStatus(ctx context.Context, orderId int, status string) error {
	orderDb, err := dbmodel.Orders(dbmodel.OrderWhere.ID.EQ(orderId),
		qm.Load(dbmodel.OrderRels.OrderItems)).One(ctx, r.dbConn)
	if err != nil {
		return err
	}

	orderDb.Status = status
	_, err = orderDb.Update(ctx, r.dbConn, boil.Whitelist(dbmodel.OrderColumns.Status))
	return err
}
