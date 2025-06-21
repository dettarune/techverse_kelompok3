// internal/service/order.go
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"todo-go/internal/model"
	"todo-go/internal/repository"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	storeRepo   *repository.StoreRepository
	productRepo *repository.ProductRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, storeRepo *repository.StoreRepository, productRepo *repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) Create(ctx context.Context, storeID int64, req *model.CreateOrderRequest) (*model.Order, string, error) {
	store, err := s.storeRepo.GetByID(ctx, storeID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get store: %w", err)
	}

	var totalAmount float64
	for _, item := range req.Items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	itemsJSON, _ := json.Marshal(req.Items)

	order := &model.Order{
		StoreID:       storeID,
		CustomerName:  req.CustomerName,
		CustomerPhone: req.CustomerPhone,
		Items:         string(itemsJSON),
		TotalAmount:   totalAmount,
		Status:        "pending",
		Notes:         req.Notes,
	}

	if err := s.orderRepo.Save(ctx, order); err != nil {
		return nil, "", fmt.Errorf("failed to save order: %w", err)
	}

	// Generate WhatsApp message
	message := fmt.Sprintf("*Pesanan Baru #%d*\n\n", order.ID)
	message += fmt.Sprintf("Nama: %s\n", req.CustomerName)
	message += fmt.Sprintf("Telepon: %s\n\n", req.CustomerPhone)
	message += "*Detail Pesanan:*\n"

	for _, item := range req.Items {
		product, _ := s.productRepo.GetByID(ctx, item.ProductID)
		if product != nil {
			message += fmt.Sprintf("- %s x%d = Rp %.0f\n", product.Name, item.Quantity, item.Price*float64(item.Quantity))
		}
	}

	message += fmt.Sprintf("\n*Total: Rp %.0f*\n", totalAmount)
	if req.Notes != "" {
		message += fmt.Sprintf("\nCatatan: %s", req.Notes)
	}

	whatsappURL := fmt.Sprintf("https://wa.me/%s?text=%s", store.WhatsApp, url.QueryEscape(message))

	return order, whatsappURL, nil
}

func (s *OrderService) GetByStore(ctx context.Context, user *model.User) ([]*model.Order, error) {
	store, err := s.storeRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	orders, err := s.orderRepo.GetByStoreID(ctx, store.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	return orders, nil
}
