package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"todo-go/internal/model"
	"todo-go/internal/service"
	"todo-go/pkg/resp"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func NewAuthHandler(authSvc *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authSvc: authSvc,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req model.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(&req)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx := r.Context()
	if err := h.authSvc.SignUp(ctx, &req); err != nil {
		switch {
		case errors.Is(err, service.ErrUserAlreadyRegistered):
			resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to create todo: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "user successfully registered",
	})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req model.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(&req)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx := r.Context()
	accessToken, err := h.authSvc.SignIn(ctx, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to create todo: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"access_token": accessToken,
	})
}
