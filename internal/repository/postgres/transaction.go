package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"time"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository"
	"traive-engineering-challenge/internal/repository/filter"
	"traive-engineering-challenge/internal/repository/models"
	"traive-engineering-challenge/internal/repository/models/mappers"
)

const TransactionModelTableExpr = "transactions"

type contextKey string

const (
	PageKey     contextKey = "page"
	PageSizeKey contextKey = "pageSize"
)

func (r *Repository) CreateTransaction(ctx context.Context, transaction domain.Transaction) (*domain.Transaction, error) {

	transactionModel, err := mappers.ConvertTransactionDomainToModel(transaction)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	transactionModel.CreatedAt = now

	_, err = r.db.NewInsert().Model(transactionModel).Exec(ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, repository.NewUniqueIndexError(pgErr.Detail)
		}
		return nil, err
	}

	transactionRecordCreated, err := mappers.ConvertTransactionModelToDomain(*transactionModel)
	if err != nil {
		return nil, err

	}
	return transactionRecordCreated, err
}

// ListTransactions retrieves a list of transactions from the database
// It accepts a context and a list of filter options
// It returns a list of transactions and an error
func (r *Repository) ListTransactions(ctx context.Context, filters ...filter.Options) ([]domain.Transaction, error) {
	var transactionModel []*models.Transaction

	transactionsFilter := &filter.TransactionFilter{
		Query: r.db.NewSelect().
			ModelTableExpr(TransactionModelTableExpr).
			Model(&transactionModel),
	}

	for _, opt := range filters {
		opt(transactionsFilter)
	}

	// Retrieve the page and pageSize from context
	page, ok := ctx.Value(PageKey).(int)
	if !ok {
		page = 1
	}
	pageSize, ok := ctx.Value(PageSizeKey).(int)
	if !ok {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	err := transactionsFilter.Query.Offset(offset).Limit(pageSize).Scan(ctx)
	if err != nil {
		return nil, errors.New("failed to list transactions")
	}

	if len(transactionModel) == 0 {
		return []domain.Transaction{}, nil
	}

	result := mappers.ConvertTransactionToDomainList(transactionModel)

	return result, nil
}
