// internal/handler/website.go
package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"todo-go/internal/model"
	"todo-go/internal/service"
	"todo-go/pkg/qr"
	"todo-go/pkg/resp"

	"github.com/go-playground/validator/v10"
)

type WebsiteHandler struct {
	websiteSvc *service.WebsiteService
	qrSvc      *qr.Service
}

func NewWebsiteHandler(websiteSvc *service.WebsiteService, qrSvc *qr.Service) *WebsiteHandler {
	return &WebsiteHandler{
		websiteSvc: websiteSvc,
		qrSvc:      qrSvc,
	}
}

func (h *WebsiteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateWebsiteRequest
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

	website, err := h.websiteSvc.Create(ctx, user, &req)
	if err != nil {
		log.Printf("failed to create website: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "website successfully created",
		"data":    website,
	})
}

func (h *WebsiteHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	website, err := h.websiteSvc.GetByUser(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrWebsiteNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to get website: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"data": website,
	})
}

func (h *WebsiteHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateWebsiteRequest
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

	website, err := h.websiteSvc.Update(ctx, user, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrWebsiteNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		default:
			log.Printf("failed to update website: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "website successfully updated",
		"data":    website,
	})
}

func (h *WebsiteHandler) GenerateQR(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*model.User)

	website, err := h.websiteSvc.GetByUser(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrWebsiteNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": "website not found, please create website first",
			})
			return
		default:
			log.Printf("failed to get website: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	// Generate catalog URL
	host := r.Header.Get("Host")
	if host == "" {
		host = "localhost:8080"
	}
	
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	catalogURL := fmt.Sprintf("%s://%s/catalog/%s", scheme, host, website.Domain)

	// Generate QR code
	qrCode, err := h.qrSvc.GenerateQR(catalogURL)
	if err != nil {
		log.Printf("failed to generate QR code: %s", err.Error())
		resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "failed to generate QR code",
		})
		return
	}

	// Set response headers for image
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"qr-catalog-%s.png\"", website.Domain))
	w.WriteHeader(http.StatusOK)
	w.Write(qrCode)
}

func (h *WebsiteHandler) GetCatalog(w http.ResponseWriter, r *http.Request) {
	domain := r.PathValue("domain")
	if domain == "" {
		resp.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error": "domain is required",
		})
		return
	}

	ctx := r.Context()

	catalog, err := h.websiteSvc.GetCatalog(ctx, domain)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrWebsiteNotFound):
			resp.WriteJSON(w, http.StatusNotFound, map[string]any{
				"error": "catalog not found",
				"message": "The requested store catalog does not exist or is not published",
			})
			return
		default:
			log.Printf("failed to get catalog: %s", err.Error())
			resp.WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
			return
		}
	}

	resp.WriteJSON(w, http.StatusOK, map[string]any{
		"store":    catalog.Store,
		"products": catalog.Products,
		"count":    len(catalog.Products),
	})
}