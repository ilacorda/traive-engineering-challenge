package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	t.Run("A valid Transaction Domain", func(t *testing.T) {
		id := uuid.New()
		userID := uuid.New()
		origin := "desktop-web"
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
			Origin:          "desktop-web",
			TransactionType: TransactionType(3),
			Amount:          int64(1000),
			CreatedAt:       time.Now(),
		}

		valid := IsValidTransactionType(transaction.TransactionType)

		assert.False(t, valid, "TransactionType should be invalid")
	})
}
