// internal/handler/order.go
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo-go/internal/model"
	"todo-go/internal/service"
	"todo-go/pkg/resp"

	"github.com/go-playground/validator/v10"
)

type OrderHandler struct {
	orderSvc *service.OrderService
}

func NewOrderHandler(orderSvc *service.OrderService) *OrderHandler {
	return &OrderHandler{orderSvc: orderSvc}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	storeID, err := strconv.Atoi(r.PathValue("storeId"))
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": "invalid store id",
		})
		return
	}

	var req model.CreateOrderRequest
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
	order, whatsappURL, err := h.orderSvc.Create(ctx, int64(storeID), &req)
	if err != nil {
		log.Printf("failed to create order: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message":      "order successfully created",
		"order":        order,
		"whatsapp_url": whatsappURL,
		"instructions": "Click the WhatsApp URL to send your order directly to the store",
	})
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	orders, err := h.orderSvc.GetByStore(ctx, user)
	if err != nil {
		log.Printf("failed to get orders: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"data":  orders,
		"count": len(orders),
	})
}
