// internal/model/store.go
package model

import "time"

type Store struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Logo        string    `json:"logo"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	WhatsApp    string    `json:"whatsapp"`
	UserID      int64     `json:"user_id" gorm:"index"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateStoreRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	WhatsApp    string `json:"whatsapp" validate:"required"`
}

type UpdateStoreRequest struct {
	ID          int64
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	WhatsApp    string `json:"whatsapp" validate:"required"`
	IsActive    bool   `json:"is_active"`
}