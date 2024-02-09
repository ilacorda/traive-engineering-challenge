package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionType int32

const (
	TransactionTypeUnspecified TransactionType = iota
	TransactionTypeCredit
	TransactionTypeDebit
)

type Transaction struct {
	ID              uuid.UUID `bun:",pk,notnull,type:uuid"`
	UserID          uuid.UUID `bun:",notnull,type:uuid"`
	Origin          string
	TransactionType string
	Amount          int64
	CreatedAt       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
