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

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Initiate database connection
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)
	if err != nil {
		log.Fatalf("failed to opening db conn: %s", err.Error())
	}

	// Run auto migration
	err = db.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		log.Fatalf("failed to run db migration: %s", err.Error())
	}

	jwtSvc := jwt.NewService(os.Getenv("APP_SECRET"))
	userRepo := repository.NewUserRepository(db)
	middSvc := middleware.NewService(jwtSvc, userRepo)

	authSvc := service.NewAuthService(userRepo, jwtSvc)
	authHandler := handler.NewAuthHandler(authSvc)

	todoRepo := repository.NewTodoRepository(db)
	todoSvc := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoSvc)

	r := http.NewServeMux()
	r.Handle("POST /api/v1/auth/signup", http.HandlerFunc(authHandler.SignUp))
	r.Handle("POST /api/v1/auth/signin", http.HandlerFunc(authHandler.SignIn))
	r.Handle("POST /api/v1/todos", middSvc.JWT(http.HandlerFunc(todoHandler.Create)))
	r.Handle("GET /api/v1/todos", middSvc.JWT(http.HandlerFunc(todoHandler.GetAllByUser)))
	r.Handle("GET /api/v1/todos/{id}", middSvc.JWT(http.HandlerFunc(todoHandler.GetByID)))
	r.Handle("PUT /api/v1/todos/{id}", middSvc.JWT(http.HandlerFunc(todoHandler.Update)))
	r.Handle("DELETE /api/v1/todos/{id}", middSvc.JWT(http.HandlerFunc(todoHandler.Delete)))

	// Apply cors and log middleware
	corsHandler := handlers.CORS()(r)
	h := handlers.LoggingHandler(os.Stdout, corsHandler)

	log.Println("server serving on port 8080")

	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Printf("could not listen server: %s", err.Error())
	}
}
