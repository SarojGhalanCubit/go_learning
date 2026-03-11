package main

import (
	"context"
	"go-minimal/internal/config"
	"go-minimal/internal/handler"
	myMw "go-minimal/internal/middleware"
	"go-minimal/internal/repository"
	"go-minimal/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware" // optional for logging
)


func main() {

	// Load ENV
	config.LoadEnv()
	db := config.ConnectDB()
	defer db.Close(context.Background())

	// Initialize repo, service, handler
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	// Initialize chi router
	r := chi.NewRouter()

	r.Use(chiMiddleware.Recoverer) // Recovers from panics
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)   // logs each request
	r.Use(chiMiddleware.Timeout(60 * 1e9)) // 60s timeout

	// Public routes
	r.Post("/register", handler.CreateUser)
	r.Post("/login", handler.Login)


	// Protected Routes
	r.Group(func(r chi.Router){
		r.Use(myMw.AuthMiddleware) // all routes here require auth

		r.Get("/users", handler.GetUsers)
		r.Get("/users/{id}", handler.GetUserByID)
	})
	
	log.Println("Server Running on http://localhost:8082")


	// http.ListenAndServe(":8080", mux)

	port := config.GetPort() // change port
	err := http.ListenAndServe(port, r)
	if err != nil {
    		log.Println("Server failed to start:", err)
	}

}
