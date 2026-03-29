package main

import (
	"context"
	"go-minimal/internal/app"
	"go-minimal/internal/app/router"
	"go-minimal/internal/config"
	"go-minimal/internal/constants"
	"log"
	"net/http"
)

func main() {

	// Load ENV
	config.LoadEnv()
	db := config.ConnectDB()
	defer db.Close(context.Background())
	constants.LoadRoles()

	// Creating Application Container
	deps := app.NewApp(db)

	// initializingglobal router
	r := router.NewRouter(deps)

	port := config.GetPort() // change port

	log.Println("Server Running on http://localhost:", port)
	err := http.ListenAndServe(port, r.Handler())
	if err != nil {
		log.Println("Server failed to start:", err)
	}

}
