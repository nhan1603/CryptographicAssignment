package orders

import (
	"context"
	"errors"
	"time"

	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
)

func (c impl) CreateOrder(ctx context.Context, order model.Order) (int, error) {
	if order.UserID <= 0 || len(order.Items) == 0 {
		return 0, errors.New("invalid order data")
	}

	// Calculate total amount and validate items
	var totalAmount float64
	for indx, item := range order.Items {
		menuItem, err := c.repo.Menu().GetByID(ctx, int(item.MenuItemID))
		if err != nil {
			return 0, err
		}

		if !menuItem.IsAvailable {
			return 0, errors.New("menu item not available")
		}

		item.UnitPrice = menuItem.Price
		item.Subtotal = menuItem.Price * float64(item.Quantity)
		totalAmount += item.Subtotal
		order.Items[indx] = item
	}

	order.TotalAmount = totalAmount
	order.Status = model.OrderStatusPending
	order.CreatedAt = time.Now()

	return c.repo.Order().Create(ctx, order)
}

func (c impl) GetOrderByID(ctx context.Context, id int) (model.Order, error) {
	order, err := c.repo.Order().GetByID(ctx, id)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (c impl) GetUserOrders(ctx context.Context, userID int) ([]model.Order, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return c.repo.Order().GetByUserID(ctx, userID)
}

func (c impl) UpdateOrderStatus(ctx context.Context, id int, status string) error {
	if !isValidOrderStatus(status) {
		return errors.New("invalid order status")
	}

	_, err := c.repo.Order().GetByID(ctx, id)
	if err != nil {
		return err
	}

	return c.repo.Order().UpdateStatus(ctx, id, status)
}

func isValidOrderStatus(status string) bool {
	validStatuses := []string{
		model.OrderStatusPending,
		model.OrderStatusPaid,
		model.OrderStatusPreparing,
		model.OrderStatusReady,
		model.OrderStatusCompleted,
		model.OrderStatusCancelled,
	}

	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}
