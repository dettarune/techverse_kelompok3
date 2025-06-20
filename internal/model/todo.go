package model

import (
	"time"
)

type Todo struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	IsComplete bool      `json:"is_complete"`
	UserID     int64     `json:"user_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	Title      string `json:"title" validate:"required"`
	IsComplete bool   `json:"is_complete"`
}

type UpdateTodoRequest struct {
	ID         int64
	Title      string `json:"title" validate:"required"`
	IsComplete bool   `json:"is_complete"`
}
