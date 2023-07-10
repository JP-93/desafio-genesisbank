package main

import (
	"desafio-genesisbank/database"
	"desafio-genesisbank/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectionDB()
	fmt.Println("API-Cotação is running....")
	r := chi.NewRouter()
	r.Get("/exchange/{amout}/{from}/{to}/{rate}", handler.ExchangeValue)
	r.Get("/exchanges", handler.GetAllValues)
	http.ListenAndServe(":8000", r)

}
