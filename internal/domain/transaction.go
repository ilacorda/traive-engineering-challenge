package domain

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

var TransactionTypeKey = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "CREDIT TRANSACTION",
	2: "DEBIT TRANSACTION",
}

var TransactionTypeValue = map[string]int32{
	"TYPE_UNSPECIFIED":   0,
	"CREDIT TRANSACTION": 1,
	"DEBIT TRANSACTION":  2,
}

func (s TransactionType) String() string {
	return [...]string{"TYPE_UNSPECIFIED", "CREDIT TRANSACTION", "DEBIT TRANSACTION"}[s]
}

// Transaction represents a money transaction
// swagger:domain Transaction
type Transaction struct {
	ID              uuid.UUID       `json:"id" validate:"required"`
	UserID          uuid.UUID       `json:"user_id" validate:"required"`
	Origin          string          `json:"origin"`
	TransactionType TransactionType `json:"transaction_type" validate:"required,oneof=0 1 2"`
	Amount          int64           `json:"amount" validate:"required"`
	CreatedAt       time.Time       `json:"created_at"`
}

func IsValidTransactionType(tType TransactionType) bool {
	switch tType {
	case TransactionTypeUnspecified, TransactionTypeCredit, TransactionTypeDebit:
		return true
	default:
		return false
	}
}

func GetTransactionTypeName(tType TransactionType) string {
	if name, ok := TransactionTypeKey[int32(tType)]; ok {
		return name
	}
	return ""
}

func StringToTransactionType(transactionType string) (TransactionType, error) {
	switch transactionType {
	case "CREDIT TRANSACTION":
		return TransactionTypeCredit, nil
	case "DEBIT TRANSACTION":
		return TransactionTypeDebit, nil
	default:
		return TransactionTypeUnspecified, nil
	}
}
