package middleware

import (
	"context"
	"net/http"
	"todo-go/internal/repository"
	"todo-go/pkg/jwt"
	"todo-go/pkg/resp"

	jwtLib "github.com/golang-jwt/jwt"
)

type Service struct {
	jwtSvc   *jwt.Service
	userRepo *repository.UserRepository
}

func NewService(jwtSvc *jwt.Service, userRepo *repository.UserRepository) *Service {
	return &Service{
		jwtSvc:   jwtSvc,
		userRepo: userRepo,
	}
}

func (s *Service) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		ctx := r.Context()

		token, err := s.jwtSvc.ParseToken(ctx, tokenStr)
		if err != nil {
			resp.WriteJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}

		if !token.Valid {
			resp.WriteJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}

		claims := token.Claims.(jwtLib.MapClaims)
		userID := claims["user_id"].(float64)

		user, err := s.userRepo.GetByID(ctx, int64(userID))
		if err != nil {
			resp.WriteJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}

		r = r.WithContext(context.WithValue(ctx, "user", user))
		next.ServeHTTP(w, r)
	})
}
