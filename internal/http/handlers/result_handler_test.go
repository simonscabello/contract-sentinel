package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
	"github.com/simonscabello/contract-sentinel/internal/results"
)

// Mock do repositório com suporte a FindAll
type mockRepo struct {
	expected []results.Result
}

func (m *mockRepo) Save(ctx context.Context, input contracts.Contract, result contracts.ValidationResult) error {
	return nil
}

func (m *mockRepo) FindAll(ctx context.Context, query results.QueryParams) ([]results.Result, error) {
	return m.expected, nil
}

func TestGetResults_Success(t *testing.T) {
	mockData := []results.Result{
		{
			ContractPath: "./dummy.json",
			ProviderURL:  "http://localhost:3000",
			Consumer:     "despesas",
			Provider:     "vexpenses-id",
			Version:      "1.0.0",
			Success:      true,
			Output:       "Tudo certo",
		},
	}

	handler := ResultsHandler{
		Repository: &mockRepo{expected: mockData},
	}

	req := httptest.NewRequest(http.MethodGet, "/results", nil)
	rr := httptest.NewRecorder()

	handler.GetResults(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("esperado status 200, obtido %d", rr.Code)
	}

	var body []results.Result
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}

	if len(body) != 1 || body[0].Consumer != "despesas" {
		t.Fatalf("resposta inesperada: %+v", body)
	}
}
