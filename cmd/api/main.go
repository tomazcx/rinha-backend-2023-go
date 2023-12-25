package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomazcx/rinha-backend-go/config"
	"github.com/tomazcx/rinha-backend-go/internal/application"
	"github.com/tomazcx/rinha-backend-go/internal/infra/db"
	"github.com/go-chi/chi/middleware"
)

func main() {

	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading the configuration file: %v", err)
	}

	err = db.ConnectToDb(conf)

	if err != nil { 
		log.Fatalf("Error connecting to the database: %v", err)
	}
	r := chi.NewRouter() 
	r.Use(middleware.Logger)
	appRouter := application.NewApplicationRouter()
	appRouter.DefineRoutes(r)

	fmt.Println("Server running at port " + conf.WebPort + "!")
	http.ListenAndServe(":"+conf.WebPort, r)
}
