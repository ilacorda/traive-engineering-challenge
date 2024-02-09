package support

import (
	"github.com/google/uuid"
	"time"
	"traive-engineering-challenge/internal/domain"
)

func ValidDomainTransactionList(transactions ...domain.Transaction) []domain.Transaction {
	return transactions
}
func ValidDomainTransaction(
	id,
	userID uuid.UUID,
	origin string,
	transactionType string,
	amount int64,
) *domain.Transaction {
	return validTransaction(id, userID, origin, &amount)
}

func validTransaction(
	id,
	userID uuid.UUID,
	origin string,
	amount *int64,
) *domain.Transaction {
	return &domain.Transaction{
		ID:              id,
		UserID:          userID,
		TransactionType: domain.TransactionTypeCredit,
		Origin:          origin,
		Amount:          *amount,
		CreatedAt:       time.Now(),
	}
}
