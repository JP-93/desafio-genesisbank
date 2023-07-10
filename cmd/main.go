package main

import (
	"desafio-genesisbank/database"
	"desafio-genesisbank/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectionDB()
	r := chi.NewRouter()
	r.Get("/exchange/{amout}/{from}/{to}/{rate}", handler.ExchangeValue)
	http.ListenAndServe(":8000", r)

}
