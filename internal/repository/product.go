// internal/repository/product.go
package repository

import (
	"context"
	"todo-go/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) GetByStoreID(ctx context.Context, storeID int64) ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.WithContext(ctx).Find(&products, "store_id = ? AND is_active = ?", storeID, true).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetByIDAndStoreID(ctx context.Context, id, storeID int64) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).First(&product, "id = ? AND store_id = ?", id, storeID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Product{}, id).Error
}
