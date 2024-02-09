package filter

import (
	"github.com/uptrace/bun"
)

const (
	Origin          string = "origin"
	TransactionType string = "transaction_type"
)

type Options func(*TransactionFilter)

type TransactionFilter struct {
	Query *bun.SelectQuery
}

func WithOrigin(origin string) Options {
	return func(f *TransactionFilter) {
		f.Query = f.Query.Where("? = ?", bun.Ident(Origin), origin)
	}
}

func WithTransactionType(transactionType string) Options {
	return func(f *TransactionFilter) {
		f.Query = f.Query.Where("? = ?", bun.Ident(TransactionType), transactionType)
	}
}
