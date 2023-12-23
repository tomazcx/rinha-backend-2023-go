package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomazcx/rinha-backend-go/config"
)

func main() {

	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading the configuration file: %v", err)
	}

	err = config.ConnectToDb(conf)

	if err != nil { 
		log.Fatalf("Error connecting to the database: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	fmt.Println("Server running at port " + conf.WebPort + "!")
	http.ListenAndServe(":"+conf.WebPort, nil)
}
