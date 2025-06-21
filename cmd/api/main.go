package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-go/internal/handler"
	"todo-go/internal/model"
	"todo-go/internal/repository"
	"todo-go/internal/service"
	"todo-go/pkg/jwt"
	"todo-go/pkg/middleware"
	"todo-go/pkg/qr"

	"github.com/gorilla/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Konfigurasi koneksi MySQL
	dbUser := "detarune"
	dbPass := "detarunism"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "todo"

	// Format DSN untuk MySQL
	// Bentuk: username:password@tcp(host:port)/dbname?parseTime=true
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to open database connection: %s", err.Error())
	}

	// Run auto migration for all models
	err = db.AutoMigrate(
		&model.User{},
		&model.Todo{},
		&model.Store{},
		&model.Product{},
		&model.Website{},
		&model.Order{},
	)
	if err != nil {
		log.Fatalf("failed to run database migration: %s", err.Error())
	}

	// Initialize core services
	jwtSvc := jwt.NewService("secretttt")
	qrSvc := qr.NewService()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	storeRepo := repository.NewStoreRepository(db)
	productRepo := repository.NewProductRepository(db)
	websiteRepo := repository.NewWebsiteRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	todoRepo := repository.NewTodoRepository(db)

	// Initialize middleware service
	middSvc := middleware.NewService(jwtSvc, userRepo)

	// Initialize business logic services
	authSvc := service.NewAuthService(userRepo, jwtSvc)
	storeSvc := service.NewStoreService(storeRepo)
	productSvc := service.NewProductService(productRepo, storeRepo)
	websiteSvc := service.NewWebsiteService(websiteRepo, storeRepo, productRepo)
	orderSvc := service.NewOrderService(orderRepo, storeRepo, productRepo)
	todoSvc := service.NewTodoService(todoRepo)

	// Initialize HTTP handlers
	authHandler := handler.NewAuthHandler(authSvc)
	storeHandler := handler.NewStoreHandler(storeSvc)
	productHandler := handler.NewProductHandler(productSvc)
	websiteHandler := handler.NewWebsiteHandler(websiteSvc, qrSvc)
	orderHandler := handler.NewOrderHandler(orderSvc)
	todoHandler := handler.NewTodoHandler(todoSvc)

	// Setup HTTP router and routes
	r := http.NewServeMux()

	// Authentication routes
	r.Handle("POST /api/v1/auth/signup", http.HandlerFunc(authHandler.SignUp))
	r.Handle("POST /api/v1/auth/signin", http.HandlerFunc(authHandler.SignIn))

	// Store management routes (protected)
	r.Handle("POST /api/v1/store", middSvc.JWT(http.HandlerFunc(storeHandler.Create)))
	r.Handle("GET /api/v1/store", middSvc.JWT(http.HandlerFunc(storeHandler.Get)))
	r.Handle("PUT /api/v1/store", middSvc.JWT(http.HandlerFunc(storeHandler.Update)))

	// Product management routes (protected)
	r.Handle("POST /api/v1/products", middSvc.JWT(http.HandlerFunc(productHandler.Create)))
	r.Handle("GET /api/v1/products", middSvc.JWT(http.HandlerFunc(productHandler.GetAll)))
	r.Handle("GET /api/v1/products/{id}", middSvc.JWT(http.HandlerFunc(productHandler.GetByID)))
	r.Handle("PUT /api/v1/products/{id}", middSvc.JWT(http.HandlerFunc(productHandler.Update)))
	r.Handle("DELETE /api/v1/products/{id}", middSvc.JWT(http.HandlerFunc(productHandler.Delete)))

	// Website builder routes (protected)
	r.Handle("POST /api/v1/website", middSvc.JWT(http.HandlerFunc(websiteHandler.Create)))
	r.Handle("GET /api/v1/website", middSvc.JWT(http.HandlerFunc(websiteHandler.Get)))
	r.Handle("PUT /api/v1/website", middSvc.JWT(http.HandlerFunc(websiteHandler.Update)))
	r.Handle("GET /api/v1/website/qr", middSvc.JWT(http.HandlerFunc(websiteHandler.GenerateQR)))

	// Public catalog route (no authentication needed)
	r.Handle("GET /catalog/{domain}", http.HandlerFunc(websiteHandler.GetCatalog))

	// Order management routes
	r.Handle("POST /api/v1/orders/{storeId}", http.HandlerFunc(orderHandler.Create))   // Public - for customers
	r.Handle("GET /api/v1/orders", middSvc.JWT(http.HandlerFunc(orderHandler.GetAll))) // Protected - for store owners

	// Todo routes (existing functionality)
	r.Handle("POST /api/v1/todos", middSvc.JWT(http.HandlerFunc(todoHandler.Create)))
	r.Handle("GET /api/v1/todos", middSvc.JWT(http.HandlerFunc(todoHandler.GetAllByUser)))
	r.Handle("GET /api/v1/todos/{id}", middSvc.JWT(http.HandlerFunc(todoHandler.GetByID)))
	r.Handle("PUT /api/v1/todos/{id}", middSvc.JWT(http.HandlerFunc(todoHandler.Update)))
	r.Handle("DELETE /api/v1/todos/{id}", middSvc.JWT(http.HandlerFunc(todoHandler.Delete)))

	// Apply CORS middleware for web access
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins in development
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
		handlers.AllowCredentials(),
	)(r)

	// Apply logging middleware
	loggedHandler := handlers.LoggingHandler(os.Stdout, corsHandler)

	// Print startup information
	log.Println("🚀 UMKM Backend API Server Starting...")
	log.Println("📍 Server running on: http://localhost:8080")
	log.Println("")
	log.Println("📋 Available API Endpoints:")
	log.Println("  Auth:")
	log.Println("    POST /api/v1/auth/signup     - Register new user")
	log.Println("    POST /api/v1/auth/signin     - Login user")
	log.Println("")
	log.Println("  Store Management:")
	log.Println("    POST /api/v1/store           - Create store profile")
	log.Println("    GET  /api/v1/store           - Get store profile")
	log.Println("    PUT  /api/v1/store           - Update store profile")
	log.Println("")
	log.Println("  Product Management:")
	log.Println("    POST   /api/v1/products      - Add new product")
	log.Println("    GET    /api/v1/products      - List all products")
	log.Println("    GET    /api/v1/products/{id} - Get product details")
	log.Println("    PUT    /api/v1/products/{id} - Update product")
	log.Println("    DELETE /api/v1/products/{id} - Delete product")
	log.Println("")
	log.Println("  Website Builder:")
	log.Println("    POST /api/v1/website         - Create website")
	log.Println("    GET  /api/v1/website         - Get website")
	log.Println("    PUT  /api/v1/website         - Update website")
	log.Println("    GET  /api/v1/website/qr      - Generate QR code")
	log.Println("")
	log.Println("  Public Access:")
	log.Println("    GET  /catalog/{domain}       - View public catalog")
	log.Println("")
	log.Println("  Order Management:")
	log.Println("    POST /api/v1/orders/{storeId} - Create order (public)")
	log.Println("    GET  /api/v1/orders          - View store orders")
	log.Println("")
	log.Println("  Todo (Legacy):")
	log.Println("    POST   /api/v1/todos         - Create todo")
	log.Println("    GET    /api/v1/todos         - List todos")
	log.Println("    GET    /api/v1/todos/{id}    - Get todo")
	log.Println("    PUT    /api/v1/todos/{id}    - Update todo")
	log.Println("    DELETE /api/v1/todos/{id}    - Delete todo")
	log.Println("")
	log.Println("🔑 Protected endpoints require 'Authorization: Bearer <token>' header")
	log.Println("📱 QR codes link to: http://localhost:8080/catalog/{domain}")
	log.Println("💬 Orders automatically generate WhatsApp URLs")
	log.Println("")

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", loggedHandler); err != nil {
		log.Fatalf("❌ Failed to start server: %s", err.Error())
	}
}
