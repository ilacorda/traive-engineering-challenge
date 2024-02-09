package service

import (
	"context"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository/filter"
)

func (t transactionService) CreateTransaction(ctx context.Context, transaction domain.Transaction) (*domain.Transaction, error) {
	result, err := t.repo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err // Return the error if there's an issue
	}
	return result, nil
}

func (t transactionService) ListTransactions(ctx context.Context, options ...filter.Options) ([]domain.Transaction, error) {
	result, err := t.repo.ListTransactions(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
