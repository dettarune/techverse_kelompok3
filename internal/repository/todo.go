package repository

import (
	"context"
	"todo-go/internal/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Save(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Save(&todo).Error
}

func (r *TodoRepository) GetByIDAndUserID(ctx context.Context, id int64, userID int64) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.WithContext(ctx).First(&todo, "id = ? AND user_id = ?", id, userID).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}

func (r *TodoRepository) GetAllByUserID(ctx context.Context, userID int64) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := r.db.WithContext(ctx).Find(&todos, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}
