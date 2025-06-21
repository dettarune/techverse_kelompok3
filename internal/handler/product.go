// internal/handler/product.go
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

type ProductHandler struct {
	productSvc *service.ProductService
}

func NewProductHandler(productSvc *service.ProductService) *ProductHandler {
	return &ProductHandler{productSvc: productSvc}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateProductRequest
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

	product, err := h.productSvc.Create(ctx, user, &req)
	if err != nil {
		log.Printf("failed to create product: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "product successfully created",
		"data":    product,
	})
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	products, err := h.productSvc.GetByStore(ctx, user)
	if err != nil {
		log.Printf("failed to get products: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"data":  products,
		"count": len(products),
	})
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": "invalid product id",
		})
		return
	}

	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	product, err := h.productSvc.GetByID(ctx, user, int64(id))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrProductNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to get product: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"data": product,
	})
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": "invalid product id",
		})
		return
	}

	var req model.UpdateProductRequest
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

	product, err := h.productSvc.Update(ctx, user, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrProductNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to update product: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "product successfully updated",
		"data":    product,
	})
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": "invalid product id",
		})
		return
	}

	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	err = h.productSvc.Delete(ctx, user, int64(id))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrProductNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to delete product: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "product successfully deleted",
	})
}
