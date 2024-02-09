package handlers

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"traive-engineering-challenge/internal/api"
)

func NewRouter(app api.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Post("/v1/transactions", CreateTransaction(app.TransactionService))
	r.Get("/v1/transactions", ListTransactions(app.TransactionService))
	return r
}
