// internal/repository/order.go
package repository

import (
	"context"
	"todo-go/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

func (r *OrderRepository) GetByID(ctx context.Context, id int64) (*model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetByStoreID(ctx context.Context, storeID int64) ([]*model.Order, error) {
	var orders []*model.Order
	err := r.db.WithContext(ctx).Find(&orders, "store_id = ?", storeID).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}