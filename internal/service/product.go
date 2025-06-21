// internal/service/product.go
package service

import (
	"context"
	"errors"
	"fmt"
	"todo-go/internal/model"
	"todo-go/internal/repository"
	"gorm.io/gorm"
)

var ErrProductNotFound = errors.New("product not found")

type ProductService struct {
	productRepo *repository.ProductRepository
	storeRepo   *repository.StoreRepository
}

func NewProductService(productRepo *repository.ProductRepository, storeRepo *repository.StoreRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		storeRepo:   storeRepo,
	}
}

func (s *ProductService) Create(ctx context.Context, user *model.User, req *model.CreateProductRequest) (*model.Product, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Image:       req.Image,
		Category:    req.Category,
		Stock:       req.Stock,
		StoreID:     store.ID,
		IsActive:    true,
	}

	if err := s.productRepo.Save(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to save product: %w", err)
	}

	return product, nil
}

func (s *ProductService) GetByStore(ctx context.Context, user *model.User) ([]*model.Product, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	products, err := s.productRepo.GetByStoreID(ctx, store.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

func (s *ProductService) GetByID(ctx context.Context, user *model.User, id int64) (*model.Product, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	product, err := s.productRepo.GetByIDAndStoreID(ctx, id, store.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}

func (s *ProductService) Update(ctx context.Context, user *model.User, req *model.UpdateProductRequest) (*model.Product, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	product, err := s.productRepo.GetByIDAndStoreID(ctx, req.ID, store.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Image = req.Image
	product.Category = req.Category
	product.Stock = req.Stock
	product.IsActive = req.IsActive

	if err := s.productRepo.Save(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return product, nil
}

func (s *ProductService) Delete(ctx context.Context, user *model.User, id int64) error {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to get store: %w", err)
	}

	product, err := s.productRepo.GetByIDAndStoreID(ctx, id, store.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return fmt.Errorf("failed to get product: %w", err)
	}

	if err := s.productRepo.Delete(ctx, product.ID); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}