package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"traive-engineering-challenge/internal/api/handlers/httperrors"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository/filter"
	"traive-engineering-challenge/internal/service"
	"traive-engineering-challenge/internal/support"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	PageKey         = "page"
	PageSizeKey     = "pageSize"
	Origin          = "origin"
	TransactionType = "transactionType"
	Message         = "message"
)

// CreateTransaction godoc
// @Summary Add a new transaction
// @Description Creates a new transaction in the system
// @tags transactions
// @Accept json
// @Produce json
// @Success 200 {array} domain.Transaction
// @Failure 500 {object} httperrors.HTTPError
// @Failure 400 {object} httperrors.HTTPError
// @Router /v1/transactions [post]
func CreateTransaction(app service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction domain.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			sendError(w, httperrors.NewHTTPError(support.ErrFailedToDecodeRequest, http.StatusBadRequest))
			return
		}

		_, err := app.CreateTransaction(r.Context(), transaction)
		if err != nil {
			sendError(w, httperrors.NewHTTPError(support.ErrFailedToCreateTransaction, http.StatusInternalServerError))
			return
		}

		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]string{Message: support.MsgTransactionCreatedSuccessfully}); err != nil {
			sendError(w, httperrors.NewHTTPError(support.ErrFailedToEncodeResponse, http.StatusInternalServerError))
		}
	}
}

// ListTransactions godoc
// @Summary List transactions
// @Description Retrieves a list of transactions based on filter criteria
// @tags transactions
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page for pagination"
// @Param origin query string false "Filter by transaction origin"
// @Param transactionType query string false "Filter by transaction type"
// @Success 200 {array} domain.Transaction
// @Failure 500 {object} httperrors.HTTPError
// @Router /v1/transactions [get]
func ListTransactions(app service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract 'page' and 'pageSize' from query parameters
		page := getQueryParamAsInt(r, PageKey, 1)
		pageSize := getQueryParamAsInt(r, PageSizeKey, 10)

		opts := extractAndBuildFilterParams(r)

		// Create a new context with the page and pageSize values
		ctxWithPagination := context.WithValue(r.Context(), PageKey, page)
		ctxWithPagination = context.WithValue(ctxWithPagination, pageSize, pageSize)

		transactions, err := app.ListTransactions(ctxWithPagination, opts...)
		if err != nil {
			sendError(w, httperrors.NewHTTPError(support.ErrFailedToRetrieveTransactions, http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(transactions); err != nil {
			sendError(w, httperrors.NewHTTPError(support.ErrFailedToEncodeResponse, http.StatusInternalServerError))
		}
	}
}

func extractAndBuildFilterParams(r *http.Request) []filter.Options {
	// Extract filter parameters
	origin := r.URL.Query().Get(Origin)
	transactionType := r.URL.Query().Get(TransactionType)

	// Create filter options based on the query parameters
	var opts []filter.Options
	if origin != "" {
		opts = append(opts, filter.WithOrigin(origin))
	}
	if transactionType != "" {
		opts = append(opts, filter.WithTransactionType(transactionType))
	}
	return opts
}

func sendError(w http.ResponseWriter, httpErr httperrors.HTTPError) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(httpErr.StatusCode)

	if err := json.NewEncoder(w).Encode(httpErr); err != nil {
		http.Error(w, "Failed to send error response", http.StatusInternalServerError)
	}
}

func getQueryParamAsInt(r *http.Request, param string, defaultVal int) int {
	valueStr := r.URL.Query().Get(param)
	if value, err := strconv.Atoi(valueStr); err == nil && value > 0 {
		return value
	}
	return defaultVal
}
