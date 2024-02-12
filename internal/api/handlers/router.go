package handlers

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
	"traive-engineering-challenge/internal/api"
)

func NewRouter(app api.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Use toHTTPHandlerFunc directly without the otelhttp prefix
	r.Post("/v1/transactions", toHTTPHandlerFunc(otelhttp.NewHandler(CreateTransaction(app.TransactionService), "CreateTransaction")))
	r.Get("/v1/transactions", toHTTPHandlerFunc(otelhttp.NewHandler(ListTransactions(app.TransactionService), "ListTransactions")))
	return r
}

// toHTTPHandlerFunc converts a http.Handler to a http.HandlerFunc
func toHTTPHandlerFunc(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}
