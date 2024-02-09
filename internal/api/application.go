//go:generate mockgen -package=mocks -source=application.go -destination=mocks/application.go . Application
package api

import (
	"traive-engineering-challenge/internal/repository"
	"traive-engineering-challenge/internal/service"
)

type Application struct {
	repository.Repository
	TransactionService service.TransactionService
}

func NewApplication(repo repository.Repository) Application {
	return Application{
		Repository:         repo,
		TransactionService: service.NewTransactionService(repo),
	}
}
