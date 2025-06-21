// internal/repository/website.go
package repository

import (
	"context"
	"todo-go/internal/model"

	"gorm.io/gorm"
)

type WebsiteRepository struct {
	db *gorm.DB
}

func NewWebsiteRepository(db *gorm.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

func (r *WebsiteRepository) Save(ctx context.Context, website *model.Website) error {
	return r.db.WithContext(ctx).Save(website).Error
}

func (r *WebsiteRepository) GetByStoreID(ctx context.Context, storeID int64) (*model.Website, error) {
	var website model.Website
	err := r.db.WithContext(ctx).First(&website, "store_id = ?", storeID).Error
	if err != nil {
		return nil, err
	}
	return &website, nil
}

func (r *WebsiteRepository) GetByDomain(ctx context.Context, domain string) (*model.Website, error) {
	var website model.Website
	err := r.db.WithContext(ctx).First(&website, "domain = ? AND is_published = ?", domain, true).Error
	if err != nil {
		return nil, err
	}
	return &website, nil
}
