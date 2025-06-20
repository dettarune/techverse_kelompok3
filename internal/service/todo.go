package service

import (
	"context"
	"errors"
	"fmt"
	"todo-go/internal/model"
	"todo-go/internal/repository"

	"gorm.io/gorm"
)

var ErrTodoNotFound = errors.New("todo not found")

type TodoService struct {
	todoRepo *repository.TodoRepository
}

func NewTodoService(todoRepo *repository.TodoRepository) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (s *TodoService) Create(ctx context.Context, user *model.User, req *model.CreateTodoRequest) (*model.Todo, error) {
	// Init todo instance
	todo := &model.Todo{
		Title:      req.Title,
		IsComplete: req.IsComplete,
		UserID:     user.ID,
	}

	// Save todo to the database
	if err := s.todoRepo.Save(ctx, todo); err != nil {
		return nil, fmt.Errorf("failed to save todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) GetByID(ctx context.Context, user *model.User, id int64) (*model.Todo, error) {
	// Get existing todo
	todo, err := s.todoRepo.GetByIDAndUserID(ctx, id, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTodoNotFound
		}

		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) Update(ctx context.Context, user *model.User, req *model.UpdateTodoRequest) (*model.Todo, error) {
	// Get existing todo
	todo, err := s.todoRepo.GetByIDAndUserID(ctx, req.ID, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTodoNotFound
		}

		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	// Save todo to the database
	todo.Title = req.Title
	todo.IsComplete = req.IsComplete
	if err := s.todoRepo.Save(ctx, todo); err != nil {
		return nil, fmt.Errorf("failed to save todo: %w", err)
	}

	return todo, nil
}

func (s *TodoService) Delete(ctx context.Context, user *model.User, id int64) error {
	// Get existing todo
	todo, err := s.todoRepo.GetByIDAndUserID(ctx, id, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTodoNotFound
		}

		return fmt.Errorf("failed to get todo: %w", err)
	}

	// Save todo to the database
	if err := s.todoRepo.Delete(ctx, todo.ID); err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	return nil
}

func (s *TodoService) GetAllByUser(ctx context.Context, user *model.User) ([]*model.Todo, error) {
	// Get all todos by user
	todos, err := s.todoRepo.GetAllByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all todos by user: %w", err)
	}

	return todos, nil
}
