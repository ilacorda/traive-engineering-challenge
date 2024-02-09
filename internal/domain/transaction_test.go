package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	t.Run("A valid Transaction Domain", func(t *testing.T) {
		id := uuid.New()
		userID := uuid.New()
		origin := "web"
		transactionType := TransactionTypeCredit
		amount := int64(1000)
		createdAt := time.Now()

		transaction := Transaction{
			ID:              id,
			UserID:          userID,
			Origin:          origin,
			TransactionType: transactionType,
			Amount:          amount,
			CreatedAt:       createdAt,
		}

		assert.Equal(t, id, transaction.ID)
		assert.Equal(t, userID, transaction.UserID)
		assert.Equal(t, origin, transaction.Origin)
		assert.Equal(t, transactionType, transaction.TransactionType)
		assert.Equal(t, amount, transaction.Amount)
		assert.Equal(t, createdAt, transaction.CreatedAt)
	})

	t.Run("An invalid Transaction Type", func(t *testing.T) {
		transaction := Transaction{
			ID:              uuid.New(),
			UserID:          uuid.New(),
			Origin:          "web",
			TransactionType: TransactionType(3), // invalid transaction type
			Amount:          int64(1000),
			CreatedAt:       time.Now(),
		}

		// Check if the transaction type is invalid using the IsValidTransactionType function
		valid := IsValidTransactionType(transaction.TransactionType)

		assert.False(t, valid, "TransactionType should be invalid")
	})
}
