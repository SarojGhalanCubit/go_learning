package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware" // optional for logging
	"go-minimal/internal/config"
	"go-minimal/internal/constants"
	"go-minimal/internal/handler"
	myMw "go-minimal/internal/middleware"
	"go-minimal/internal/repository"
	"go-minimal/internal/service"
	"log"
	"net/http"
)

func main() {

	// Load ENV
	config.LoadEnv()
	db := config.ConnectDB()
	defer db.Close(context.Background())

	// Initialize repo, service, handler
	repo := repository.NewUserRepository(db)
	materialRepo := repository.NewMaterialRepository(db)

	// Initialize Service
	userService := service.NewUserService(repo)
	materialService := service.NewMaterialService(materialRepo)

	userHandler := handler.NewUserHandler(userService)
	materialHandler := handler.NewMaterialHandler(materialService)

	// Initialize chi router
	r := chi.NewRouter()

	r.Use(chiMiddleware.Recoverer) // Recovers from panics
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)            // logs each request
	r.Use(chiMiddleware.Timeout(60 * 1e9)) // 60s timeout

	// Public routes
	r.Post("/login", userHandler.Login)

	// Protected Routes
	r.Group(func(r chi.Router) {
		r.Use(myMw.AuthMiddleware) // all routes here require auth

		// Admin Only routes
		r.Group(func(r chi.Router) {
			r.Use(myMw.RequireRole(constants.Admin))
			r.Put("/user/{id}/update", userHandler.UpdateUser)
			r.Delete("/user/{id}/delete", userHandler.DeleteUser)

		})

		// Admin And Manager Routes
		r.Group(func(r chi.Router) {
			r.Use(myMw.RequireRole(constants.Admin, constants.Manager))

			// user endpoints
			r.Get("/users/getAll", userHandler.GetUsers)
			r.Get("/user/{id}/getById", userHandler.GetUserByID)
			r.Post("/user/create", userHandler.CreateUser)

			// Material endpoints

			r.Get("/materials/getAll", materialHandler.GetAllMaterial)
		})

	})

	log.Println("Server Running on http://localhost:8082")

	port := config.GetPort() // change port
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Println("Server failed to start:", err)
	}

}
