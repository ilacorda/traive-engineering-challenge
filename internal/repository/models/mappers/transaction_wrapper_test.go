package mappers

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	repository "traive-engineering-challenge/internal/support"

	"github.com/google/uuid"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository/models"
)

func TestConvertTransactionDomainToModel(t *testing.T) {
	t.Parallel()
	validID := uuid.New()
	validUserID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name          string
		transaction   domain.Transaction
		expectedModel *models.Transaction
		expectedError error
	}{
		{
			name: "it should convert a valid transaction domain to a model transaction",
			transaction: domain.Transaction{
				ID:              validID,
				UserID:          validUserID,
				Origin:          repository.DesktopWeb,
				TransactionType: domain.TransactionTypeCredit,
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedModel: &models.Transaction{
				ID:              validID,
				UserID:          validUserID,
				Origin:          repository.DesktopWeb,
				TransactionType: string(models.TransactionTypeCredit),
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedError: nil,
		},
		{
			name: "it should return an error for an invalid transaction ID",
			transaction: domain.Transaction{
				ID:              uuid.Nil,
				UserID:          uuid.New(),
				Origin:          repository.DesktopWeb,
				TransactionType: domain.TransactionTypeCredit,
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedModel: nil,
			expectedError: errors.New("invalid transaction ID"),
		},
		{
			name: "it should return an error for an invalid user ID",
			transaction: domain.Transaction{
				ID:              uuid.New(),
				UserID:          uuid.Nil,
				Origin:          repository.DesktopWeb,
				TransactionType: domain.TransactionTypeCredit,
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedModel: nil,
			expectedError: errors.New("invalid user ID"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model, err := ConvertTransactionDomainToModel(tc.transaction)
			require.NoError(t, err)

			if err != nil {
				if tc.expectedError == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("unexpected error: got %v, want %v", err, tc.expectedError)
				} else if tc.expectedError != nil {
					if diff := cmp.Diff(tc.expectedModel, model, cmp.Comparer(generateTimeComparer())); diff != "" {
						t.Errorf("models differ (-expectedModel +model):\n%s", diff)
					}
				}
			}
		})
	}
}

func TestConvertTransactionModelToDomain(t *testing.T) {
	t.Parallel()

	validID := uuid.New()
	validUserID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name             string
		transactionModel models.Transaction
		expectedDomain   domain.Transaction
		expectedError    error
	}{
		{
			name: "it should convert a valid transaction model to a domain transaction",
			transactionModel: models.Transaction{
				ID:              validID,
				UserID:          validUserID,
				Origin:          repository.MobileAndroid,
				TransactionType: string(models.TransactionTypeDebit),
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedDomain: domain.Transaction{
				ID:              validID,
				UserID:          validUserID,
				Origin:          repository.MobileAndroid,
				TransactionType: domain.TransactionTypeDebit,
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedError: nil,
		},
		{
			name: "it should return an error for an invalid transaction type",
			transactionModel: models.Transaction{
				ID:              validID,
				UserID:          validUserID,
				Origin:          repository.MobileAndroid,
				TransactionType: "invalid",
				Amount:          1000,
				CreatedAt:       now,
			},
			expectedDomain: domain.Transaction{},
			expectedError:  errors.New("invalid transaction type"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transactionDomain, err := ConvertTransactionModelToDomain(tc.transactionModel)
			require.NoError(t, err)

			if err != nil {
				if tc.expectedError == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("unexpected error: got %v, want %v", err, tc.expectedError)
				} else if tc.expectedError != nil {
					if diff := cmp.Diff(tc.expectedDomain, transactionDomain, cmp.Comparer(generateTimeComparer())); diff != "" {
						t.Errorf("domains differ (-expectedDomain +domain):\n%s", diff)
					}
				}
			}
		})
	}
}

func generateTimeComparer() cmp.Option {
	timeComparer := cmp.Comparer(func(x, y time.Time) bool {
		diff := x.Sub(y)
		if diff < 0 {
			diff = -diff
		}
		allowedGap := time.Second
		return diff <= allowedGap
	})
	return timeComparer
}
