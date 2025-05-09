package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
)

// Mock service
type mockValidationService struct {
	result contracts.ValidationResult
	err    error
}

func (m *mockValidationService) ValidateContract(c contracts.Contract) (contracts.ValidationResult, error) {
	return m.result, m.err
}

// Mock repository
type mockResultsRepo struct{}

func (m *mockResultsRepo) Save(ctx context.Context, input contracts.Contract, result contracts.ValidationResult) error {
	return nil
}

func TestValidateContract_WithMetadata(t *testing.T) {
	mock := &mockValidationService{
		result: contracts.ValidationResult{
			Success: true,
			Output:  "OK",
		},
	}

	handler := ContractHandler{
		Service:    mock,
		Repository: &mockResultsRepo{},
	}

	body := contractRequest{
		Path:        "./dummy.json",
		ProviderURL: "http://localhost:3000",
		Consumer:    "despesas",
		Provider:    "vexpenses-id",
		Version:     "1.0.0",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/contracts/test", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ValidateContract(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("esperado status 200, obtido %d", rr.Code)
	}
}
