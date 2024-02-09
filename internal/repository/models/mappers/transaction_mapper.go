package mappers

import (
	"fmt"
	"github.com/google/uuid"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository/models"
)

// ConvertTransactionDomainToModel converts a domain.Transaction to a models.Transaction.
func ConvertTransactionDomainToModel(transaction domain.Transaction) (*models.Transaction, error) {
	id, err := ValidateUUID(transaction.ID.String())
	if err != nil {
		return nil, fmt.Errorf("invalid transaction ID: %w", err)
	}

	userID, err := ValidateUUID(transaction.UserID.String())
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	return &models.Transaction{
		ID:              id,
		UserID:          userID,
		Origin:          transaction.Origin,
		TransactionType: transaction.TransactionType.String(),
		Amount:          transaction.Amount,
		CreatedAt:       transaction.CreatedAt,
	}, nil
}

func ConvertTransactionModelToDomain(transactionModel models.Transaction) (*domain.Transaction, error) {
	transactionType, err := domain.StringToTransactionType(transactionModel.TransactionType)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction type: %w", err)
	}

	return &domain.Transaction{
		ID:              transactionModel.ID,
		UserID:          transactionModel.UserID,
		Origin:          transactionModel.Origin,
		TransactionType: domain.TransactionType(transactionType),
		Amount:          transactionModel.Amount,
		CreatedAt:       transactionModel.CreatedAt,
	}, nil
}

func ConvertTransactionToDomainList(models []*models.Transaction) []domain.Transaction {
	var domainList []domain.Transaction
	for _, model := range models {
		toDomain, _ := ConvertTransactionModelToDomain(*model)

		domainList = append(domainList, *toDomain)
	}

	return domainList
}

// ValidateUUID TODO [Improvement - code organisation] move this to a utility package
// ValidateUUID validates a UUID string and returns a UUID object.
func ValidateUUID(value string) (uuid.UUID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
