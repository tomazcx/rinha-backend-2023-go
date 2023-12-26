package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/tomazcx/rinha-backend-go/config"
	"github.com/tomazcx/rinha-backend-go/internal/application"
	"github.com/tomazcx/rinha-backend-go/internal/infra/db"
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
	uuid.EnableRandPool()

	r := chi.NewRouter() 

	appRouter := application.NewApplicationRouter()
	appRouter.DefineRoutes(r)

	fmt.Println("Server running at port " + conf.WebPort + "!")
	http.ListenAndServe(":"+conf.WebPort, r)
}
