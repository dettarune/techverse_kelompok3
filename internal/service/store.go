// internal/service/store.go
package service

import (
	"context"
	"errors"
	"fmt"
	"todo-go/internal/model"
	"todo-go/internal/repository"

	"gorm.io/gorm"
)

var ErrStoreNotFound = errors.New("store not found")

type StoreService struct {
	storeRepo *repository.StoreRepository
}

func NewStoreService(storeRepo *repository.StoreRepository) *StoreService {
	return &StoreService{storeRepo: storeRepo}
}

func (s *StoreService) Create(ctx context.Context, user *model.User, req *model.CreateStoreRequest) (*model.Store, error) {
	store := &model.Store{
		Name:        req.Name,
		Description: req.Description,
		Logo:        req.Logo,
		Address:     req.Address,
		Phone:       req.Phone,
		WhatsApp:    req.WhatsApp,
		UserID:      user.ID,
		IsActive:    true,
	}

	if err := s.storeRepo.Save(ctx, store); err != nil {
		return nil, fmt.Errorf("failed to save store: %w", err)
	}

	return store, nil
}

func (s *StoreService) GetByUser(ctx context.Context, user *model.User) (*model.Store, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStoreNotFound
		}
		return nil, fmt.Errorf("failed to get store: %w", err)
	}
	return store, nil
}

func (s *StoreService) Update(ctx context.Context, user *model.User, req *model.UpdateStoreRequest) (*model.Store, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStoreNotFound
		}
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	store.Name = req.Name
	store.Description = req.Description
	store.Logo = req.Logo
	store.Address = req.Address
	store.Phone = req.Phone
	store.WhatsApp = req.WhatsApp
	store.IsActive = req.IsActive

	if err := s.storeRepo.Save(ctx, store); err != nil {
		return nil, fmt.Errorf("failed to update store: %w", err)
	}

	return store, nil
}
