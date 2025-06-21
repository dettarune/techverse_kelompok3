package model

import "time"

type Order struct {
	ID            int64     `json:"id"`
	StoreID       int64     `json:"store_id" gorm:"index"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	Items         string    `json:"items"` // JSON string of ordered items
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"` // pending, confirmed, completed, cancelled
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
	Items         []OrderItem `json:"items" validate:"required,min=1"`
	CustomerName  string      `json:"customer_name" validate:"required"`
	CustomerPhone string      `json:"customer_phone" validate:"required"`
	Notes         string      `json:"notes"`
}

type OrderItem struct {
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
