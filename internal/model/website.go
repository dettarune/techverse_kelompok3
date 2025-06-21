// internal/model/website.go
package model

import "time"

type Website struct {
	ID          int64     `json:"id"`
	StoreID     int64     `json:"store_id" gorm:"index"`
	Template    string    `json:"template"`
	CustomCSS   string    `json:"custom_css"`
	CustomHTML  string    `json:"custom_html"`
	Domain      string    `json:"domain"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateWebsiteRequest struct {
	Template   string `json:"template" validate:"required"`
	CustomCSS  string `json:"custom_css"`
	CustomHTML string `json:"custom_html"`
	Domain     string `json:"domain"`
}

type UpdateWebsiteRequest struct {
	ID          int64
	Template    string `json:"template" validate:"required"`
	CustomCSS   string `json:"custom_css"`
	CustomHTML  string `json:"custom_html"`
	Domain      string `json:"domain"`
	IsPublished bool   `json:"is_published"`
}