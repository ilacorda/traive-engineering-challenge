//go:generate mockgen -package=mocks -source=service.go -destination=mocks/service.go . ProviderService
package service

import (
	"context"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository"
	"traive-engineering-challenge/internal/repository/filter"
)

type transactionService struct {
	repo repository.Repository
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (*domain.Transaction, error)
	ListTransactions(ctx context.Context, options ...filter.Options) ([]domain.Transaction, error)
}

func NewTransactionService(repo repository.Repository) TransactionService {
	return transactionService{
		repo: repo,
	}
}
