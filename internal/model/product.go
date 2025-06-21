package model

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Category    string    `json:"category"`
	Stock       int       `json:"stock"`
	StoreID     int64     `json:"store_id" gorm:"index"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock" validate:"min=0"`
}

type UpdateProductRequest struct {
	ID          int64
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock" validate:"min=0"`
	IsActive    bool    `json:"is_active"`
}