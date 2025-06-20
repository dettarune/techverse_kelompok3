package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo-go/internal/model"
	"todo-go/internal/repository"
	"todo-go/pkg/jwt"

	"gorm.io/gorm"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrInvalidCredentials    = errors.New("invalid email or password")
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtSvc   *jwt.Service
}

func NewAuthService(userRepo *repository.UserRepository, jwtSvc *jwt.Service) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtSvc:   jwtSvc,
	}
}

func (s *AuthService) SignUp(ctx context.Context, req *model.SignUpRequest) error {
	// Check existing user
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to get user by email: %w", err)
	}

	if user != nil {
		return ErrUserAlreadyRegistered
	}

	// Initiate user instance
	user = &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// Generate password hash
	if err := user.GeneratePassword(req.Password); err != nil {
		return fmt.Errorf("failed to generate password: %w", err)
	}

	// Save user to the database
	if err := s.userRepo.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, req *model.SignInRequest) (string, error) {
	// Get existing user
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("failed to get user by email: %w", err)
	}

	if user == nil {
		return "", ErrInvalidCredentials
	}

	// Validate password
	if !user.ValidatePassword(req.Password) {
		return "", ErrInvalidCredentials
	}

	// Generate access token
	tokenExp := time.Now().Add(24 * time.Hour).Unix()
	tokenData := map[string]any{"user_id": user.ID}

	accessToken, err := s.jwtSvc.GenerateToken(ctx, tokenData, tokenExp)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}

	return accessToken, nil
}
