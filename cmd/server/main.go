
package main

import (
	"log"
	"context"
	"net/http"
	"go-minimal/internal/handler"
	"go-minimal/internal/middleware"
	"go-minimal/internal/repository"
	"go-minimal/internal/service"
	"go-minimal/internal/config"
)


func main() {
	config.LoadEnv()
	db := config.ConnectDB()
	defer db.Close(context.Background())

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)



	mux := http.NewServeMux()
	

	log.Println("Server Running on http://localhost:8082")


	mux.Handle("/users", middleware.Logging(http.HandlerFunc(handler.GetUsers)))
	mux.Handle("/users/create", middleware.Logging(http.HandlerFunc(handler.CreateUser)))

	// http.ListenAndServe(":8080", mux)

	port := config.GetPort() // change port
	err := http.ListenAndServe(port, mux)
	if err != nil {
    		log.Println("Server failed to start:", err)
	}

}
