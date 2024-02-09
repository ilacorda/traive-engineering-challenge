//go:generate mockgen -package=mocks -source=repository.go -destination=mocks/repository.go . Repository

package repository

import (
	"context"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository/filter"
)

type Repository interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (*domain.Transaction, error)
	ListTransactions(ctx context.Context, filters ...filter.Options) ([]domain.Transaction, error)
}
