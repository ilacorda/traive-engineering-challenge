package postgres

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/support"
)

var (
	transactionSchema = []string{"id", "user_id", "origin", "transaction_type", "amount", "created_at"}

	transactionIDOne = "73b2228a-be4a-43dd-8c07-4668e59da688"
	transactionIDTwo = "f3b2228a-be4a-43dd-8c07-4668e59da688"

	buildPopulatedTransactions = func() *sqlmock.Rows {
		return sqlmock.NewRows(transactionSchema).
			AddRow(transactionIDOne, uuid.NewString(), "desktop-web", 2, time.Now()).
			AddRow(transactionIDTwo, uuid.NewString(), "web", 1, 1000, time.Now())
	}
)

const (
	InsertTransactionQuery = `^INSERT INTO "transactions"`
	GetTransactionsQuery   = `^SELECT FROM "transactions"`
)

type queryMock func(sqlmock.Sqlmock)

func TestRepository_CreateTransaction(t *testing.T) {
	t.Parallel()

	transaction := domain.Transaction{
		ID:              uuid.New(),
		UserID:          uuid.New(),
		Origin:          support.DesktopWeb,
		TransactionType: domain.TransactionTypeCredit,
		Amount:          1000,
		CreatedAt:       time.Now(),
	}

	testData := map[string]struct {
		setupMocks       func(sqlmock.Sqlmock)
		wantErr          bool
		inputTransaction domain.Transaction
	}{
		"happy path - creates new transaction": {
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(InsertTransactionQuery).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputTransaction: transaction,
			wantErr:          false,
		},
		"failure - insert provider fails": {
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(InsertTransactionQuery).
					WillReturnError(fmt.Errorf("insert operation failed"))
			},
			inputTransaction: transaction,
			wantErr:          true,
		},
	}

	for name, tc := range testData {
		t.Run(name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			require.NoError(t, err)

			repo, err := NewRepository(db)
			require.NoError(t, err)

			tc.setupMocks(mock)

			_, err = repo.CreateTransaction(context.Background(), tc.inputTransaction)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// Helper functions

func expectationMet(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
