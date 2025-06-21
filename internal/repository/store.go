// internal/repository/store.go
package repository

import (
	"context"
	"todo-go/internal/model"
	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) Save(ctx context.Context, store *model.Store) error {
	return r.db.WithContext(ctx).Save(store).Error
}

func (r *StoreRepository) GetByID(ctx context.Context, id int64) (*model.Store, error) {
	var store model.Store
	err := r.db.WithContext(ctx).First(&store, id).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepository) GetByUserID(ctx context.Context, userID int64) (*model.Store, error) {
	var store model.Store
	err := r.db.WithContext(ctx).First(&store, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Store{}, id).Error
}