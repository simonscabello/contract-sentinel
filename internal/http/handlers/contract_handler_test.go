package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
)

type mockValidationService struct {
	result contracts.ValidationResult
	err    error
}

func (m *mockValidationService) ValidateContract(c contracts.Contract) (contracts.ValidationResult, error) {
	return m.result, m.err
}

func TestValidateContract_Success(t *testing.T) {
	mock := &mockValidationService{
		result: contracts.ValidationResult{
			Success: true,
			Output:  "Contrato validado com sucesso.",
			Error:   nil,
		},
	}

	handler := ContractHandler{Service: mock}

	body := contractRequest{
		Path:        "./dummy.json",
		ProviderURL: "http://localhost:3000",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/contracts/test", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.ValidateContract(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("esperado status 200, obtido %d", rr.Code)
	}

	var response contracts.ValidationResult
	json.NewDecoder(rr.Body).Decode(&response)

	if !response.Success {
		t.Errorf("esperado success true, obtido false")
	}
}

func TestValidateContract_InvalidBody(t *testing.T) {
	mock := &mockValidationService{}

	handler := ContractHandler{Service: mock}

	req := httptest.NewRequest(http.MethodPost, "/contracts/test", bytes.NewReader([]byte("invalid-json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ValidateContract(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("esperado 400 Bad Request, obtido %d", rr.Code)
	}
}
