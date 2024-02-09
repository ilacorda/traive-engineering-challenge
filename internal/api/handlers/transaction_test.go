package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"traive-engineering-challenge/internal/api/handlers/httperrors"
	"traive-engineering-challenge/internal/domain"
	"traive-engineering-challenge/internal/repository/filter"
	"traive-engineering-challenge/internal/service/mocks"
	"traive-engineering-challenge/internal/support"
)

var endpoint = "/v1/transactions"

const internalServerErrorMsg = "internal server error"

// TODO write helper functions to reduce code duplication
// in order to check setup the request as well as check the response in both tests
func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTransactionService(ctrl)

	id := uuid.New()
	userID := uuid.New()
	origin := support.DesktopWeb
	transactionType := string(domain.TransactionTypeCredit)
	amount := int64(500)

	tests := []struct {
		name           string
		body           interface{}
		prepareService func()
		wantStatusCode int
		wantResponse   interface{}
	}{
		{
			name: "it returns created when the transaction is created successfully",
			body: support.ValidDomainTransaction(
				id,
				userID,
				origin,
				transactionType,
				amount,
			),
			prepareService: func() {
				mockService.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(&domain.Transaction{}, nil)
			},
			wantStatusCode: http.StatusCreated,
			wantResponse:   map[string]string{"message": "Transaction created successfully"},
		},
		{
			name: "it returns bad request when the request body is invalid",
			body: "invalid body",
			prepareService: func() {
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   httperrors.NewHTTPError("Failed to decode request body", http.StatusBadRequest),
		},
		{
			name: "it returns internal server error",
			body: support.ValidDomainTransaction(
				id,
				userID,
				origin,
				transactionType,
				amount,
			),
			prepareService: func() {
				mockService.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil, errors.New(internalServerErrorMsg))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   httperrors.NewHTTPError("Failed to create transaction", http.StatusInternalServerError),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareService()

			body, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := CreateTransaction(mockService)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.wantStatusCode)
			}

			var actualResponse map[string]interface{}
			if err := json.NewDecoder(rr.Body).Decode(&actualResponse); err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			expectedResponseJSON, err := json.Marshal(tc.wantResponse)
			if err != nil {
				t.Fatalf("Failed to marshal expected response: %v", err)
			}

			var expectedResponse map[string]interface{}
			if err := json.Unmarshal(expectedResponseJSON, &expectedResponse); err != nil {
				t.Fatalf("Failed to unmarshal expected response JSON: %v", err)
			}

			// TODO use cmp.Diff to compare actualResponse and expectedResponse
			if !reflect.DeepEqual(actualResponse, expectedResponse) {
				t.Errorf("handler returned unexpected body: got %v want %v", actualResponse, expectedResponse)
			}
		})
	}
}
func TestListTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTransactionService(ctrl)

	transactionOne := support.ValidDomainTransaction(
		uuid.New(),
		uuid.New(),
		support.MobileIOS,
		string(domain.TransactionTypeCredit),
		int64(500),
	)

	transactionTwo := support.ValidDomainTransaction(
		uuid.New(),
		uuid.New(),
		support.DesktopWeb,
		string(domain.TransactionTypeDebit),
		int64(500),
	)

	transactions := support.ValidDomainTransactionList(*transactionOne, *transactionTwo)

	tests := []struct {
		name           string
		queryParams    map[string]string
		prepareService func(mockSvc *mocks.MockTransactionService)
		wantStatusCode int
		wantResponse   interface{}
	}{
		{
			name: "it lists transactions successfully with default filters",
			queryParams: map[string]string{
				"page":     "1",
				"pageSize": "10",
			},
			prepareService: func(mockSvc *mocks.MockTransactionService) {
				mockSvc.EXPECT().ListTransactions(gomock.Any(), gomock.Any()).Return(transactions, nil)
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   transactions,
		},
		{
			name: "Successful listing with custom pagination and with filters",
			queryParams: map[string]string{
				"page":            "2",
				"pageSize":        "5",
				"origin":          support.DesktopWeb,
				"transactionType": string(domain.TransactionTypeCredit),
			},
			prepareService: func(mockSvc *mocks.MockTransactionService) {
				mockSvc.EXPECT().ListTransactions(gomock.Any(), gomock.Any()).Return([]domain.Transaction{*transactionOne}, nil)
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   []domain.Transaction{*transactionOne},
		},
		{
			name:        "Failed listing due to service error",
			queryParams: map[string]string{},
			prepareService: func(mockSvc *mocks.MockTransactionService) {
				mockSvc.EXPECT().ListTransactions(gomock.Any(), gomock.Any()).Return(nil, errors.New("service error"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   httperrors.NewHTTPError(support.ErrFailedToListTransactions, http.StatusInternalServerError),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareService(mockService)

			req, err := http.NewRequest("GET", "/transactions", nil)
			if err != nil {
				t.Fatal(err)
			}

			q := req.URL.Query()
			for key, value := range tc.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()

			handler := ListTransactions(mockService)
			handler.ServeHTTP(rr, req)

			filterObj := &filter.TransactionFilter{}
			for _, param := range tc.queryParams {
				switch {
				case param == "origin" && tc.queryParams[param] != "":
					filter.WithOrigin(tc.queryParams[param])(filterObj)
				case param == "transactionType" && tc.queryParams[param] != "":
					filter.WithTransactionType(tc.queryParams[param])(filterObj)
				}
			}

			if status := rr.Code; status != tc.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.wantStatusCode)
			}

			var actualResponse interface{}
			if err := json.NewDecoder(rr.Body).Decode(&actualResponse); err != nil {
				t.Fatalf("Failed to decode response body for %s: %v", tc.name, err)
			}
			expectedResponseJSON, err := json.Marshal(tc.wantResponse)
			if err != nil {
				t.Fatalf("Failed to marshal expected response for %s: %v", tc.name, err)
			}

			var expectedResponse interface{}
			if err := json.Unmarshal(expectedResponseJSON, &expectedResponse); err != nil {
				t.Fatalf("Failed to unmarshal expected response JSON for %s: %v", tc.name, err)
			}

			// TODO use cmptools to compare actualResponse and expectedResponse
			if !reflect.DeepEqual(actualResponse, expectedResponse) {
				t.Errorf("handler returned unexpected body for %s: got %v want %v", tc.name, actualResponse, expectedResponse)
			}
		})
	}
}
