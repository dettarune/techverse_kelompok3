// internal/handler/store.go
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

type StoreHandler struct {
	storeSvc *service.StoreService
}

func NewStoreHandler(storeSvc *service.StoreService) *StoreHandler {
	return &StoreHandler{storeSvc: storeSvc}
}

func (h *StoreHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateStoreRequest
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
	user := ctx.Value("user").(*model.User)

	store, err := h.storeSvc.Create(ctx, user, &req)
	if err != nil {
		log.Printf("failed to create store: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "store successfully created",
		"data":    store,
	})
}

func (h *StoreHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	store, err := h.storeSvc.GetByUser(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrStoreNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to get store: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"data": store,
	})
}

func (h *StoreHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateStoreRequest
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
	user := ctx.Value("user").(*model.User)

	store, err := h.storeSvc.Update(ctx, user, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrStoreNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to update store: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "store successfully updated",
		"data":    store,
	})
}
