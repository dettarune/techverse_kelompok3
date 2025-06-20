package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"todo-go/internal/model"
	"todo-go/internal/service"
	"todo-go/pkg/resp"

	"github.com/go-playground/validator/v10"
)

type TodoHandler struct {
	todoSvc *service.TodoService
}

func NewTodoHandler(todoSvc *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoSvc: todoSvc,
	}
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTodoRequest
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

	todo, err := h.todoSvc.Create(ctx, user, &req)
	if err != nil {
		log.Printf("failed to create todo: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	todo, err := h.todoSvc.GetByID(ctx, user, int64(id))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTodoNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to get todo: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	var req model.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(&req)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx := r.Context()
	user := ctx.Value("user").(*model.User)
	req.ID = int64(id)

	todo, err := h.todoSvc.Update(ctx, user, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTodoNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to update todo: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	err = h.todoSvc.Delete(ctx, user, int64(id))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTodoNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to delete todo: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "todo successfully deleted",
	})
}

func (h *TodoHandler) GetAllByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	todos, err := h.todoSvc.GetAllByUser(ctx, user)
	if err != nil {
		log.Printf("failed to get todo: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, todos)
}
