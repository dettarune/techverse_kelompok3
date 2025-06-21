// internal/service/website.go
package service

import (
	"context"
	"errors"
	"fmt"
	"todo-go/internal/model"
	"todo-go/internal/repository"

	"gorm.io/gorm"
)

var ErrWebsiteNotFound = errors.New("website not found")

type WebsiteService struct {
	websiteRepo *repository.WebsiteRepository
	storeRepo   *repository.StoreRepository
	productRepo *repository.ProductRepository
}

func NewWebsiteService(websiteRepo *repository.WebsiteRepository, storeRepo *repository.StoreRepository, productRepo *repository.ProductRepository) *WebsiteService {
	return &WebsiteService{
		websiteRepo: websiteRepo,
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}

func (s *WebsiteService) Create(ctx context.Context, user *model.User, req *model.CreateWebsiteRequest) (*model.Website, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	website := &model.Website{
		StoreID:     store.ID,
		Template:    req.Template,
		CustomCSS:   req.CustomCSS,
		CustomHTML:  req.CustomHTML,
		Domain:      req.Domain,
		IsPublished: false,
	}

	if err := s.websiteRepo.Save(ctx, website); err != nil {
		return nil, fmt.Errorf("failed to save website: %w", err)
	}

	return website, nil
}

func (s *WebsiteService) GetByUser(ctx context.Context, user *model.User) (*model.Website, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	website, err := s.websiteRepo.GetByStoreID(ctx, store.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWebsiteNotFound
		}
		return nil, fmt.Errorf("failed to get website: %w", err)
	}

	return website, nil
}

func (s *WebsiteService) Update(ctx context.Context, user *model.User, req *model.UpdateWebsiteRequest) (*model.Website, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	website, err := s.websiteRepo.GetByStoreID(ctx, store.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWebsiteNotFound
		}
		return nil, fmt.Errorf("failed to get website: %w", err)
	}

	website.Template = req.Template
	website.CustomCSS = req.CustomCSS
	website.CustomHTML = req.CustomHTML
	website.Domain = req.Domain
	website.IsPublished = req.IsPublished

	if err := s.websiteRepo.Save(ctx, website); err != nil {
		return nil, fmt.Errorf("failed to update website: %w", err)
	}

	return website, nil
}

type CatalogData struct {
	Store    *model.Store     `json:"store"`
	Products []*model.Product `json:"products"`
}

func (s *WebsiteService) GetCatalog(ctx context.Context, domain string) (*CatalogData, error) {
	website, err := s.websiteRepo.GetByDomain(ctx, domain)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWebsiteNotFound
		}
		return nil, fmt.Errorf("failed to get website: %w", err)
	}

	store, err := s.storeRepo.GetByID(ctx, website.StoreID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	products, err := s.productRepo.GetByStoreID(ctx, website.StoreID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return &CatalogData{
		Store:    store,
		Products: products,
	}, nil
}
