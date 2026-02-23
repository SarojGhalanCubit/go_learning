package main

import (
	"fmt"
	"net/http"
	"go-minimal/handler"
	"go-minimal/middleware"
	"go-minimal/repository"
	"go-minimal/service"
)


func main() {
	repo := repository.NewInMemoryUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	mux := http.NewServeMux()
	

	fmt.Println("Server Running on http://localhost:8080")

	mux.Handle("/users", middleware.Logging(http.HandlerFunc(handler.GetUsers)))
	mux.Handle("/users/create", middleware.Logging(http.HandlerFunc(handler.CreateUser)))

	http.ListenAndServe(":8080", mux)

}
