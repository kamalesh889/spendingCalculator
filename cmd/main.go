package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kamalesh889/spendingCalculator/api"
	"github.com/kamalesh889/spendingCalculator/config"
)

func main() {
	log.Println("Spending calculator : v1")

	// Needs to be changed should be read from the env file or something like that ( will implement later)
	cfg := &config.Config{
		DbURL: "postgres://postgres:kamalesh@localhost:5432/calculator",
		Port:  "8080",
	}

	server, err := api.NewServer(cfg)
	if err != nil {
		log.Panicln("Error in creating server:", err)
	}

	mux := api.Router(server)
	http.Handle("/", mux)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)

}
