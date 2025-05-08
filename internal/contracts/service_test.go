package contracts

import (
	"errors"
	"testing"
)

type mockValidatorAdapter struct {
	expectedResult ValidationResult
	expectedError  error
}

func (m *mockValidatorAdapter) Validate(contractPath string, providerURL string) (ValidationResult, error) {
	return m.expectedResult, m.expectedError
}

func TestContractValidationService_ValidateContract_Success(t *testing.T) {
	mock := &mockValidatorAdapter{
		expectedResult: ValidationResult{
			Success: true,
			Output:  "All tests passed.",
			Error:   nil,
		},
	}

	service := NewContractValidationService(mock)

	result, err := service.ValidateContract(Contract{
		Path:        "./dummy.json",
		ProviderURL: "http://localhost:3000",
	})

	if err != nil {
		t.Fatalf("esperado nil erro, obtido: %v", err)
	}
	if !result.Success {
		t.Error("esperado sucesso, obtido false")
	}
	if result.Output != "All tests passed." {
		t.Errorf("output inesperado: %s", result.Output)
	}
}

func TestContractValidationService_ValidateContract_Failure(t *testing.T) {
	mock := &mockValidatorAdapter{
		expectedResult: ValidationResult{
			Success: false,
			Output:  "Erro ao validar",
			Error:   errors.New("validação falhou"),
		},
		expectedError: errors.New("validação falhou"),
	}

	service := NewContractValidationService(mock)

	_, err := service.ValidateContract(Contract{
		Path:        "./dummy.json",
		ProviderURL: "http://localhost:3000",
	})

	if err == nil {
		t.Fatal("esperado erro, obtido nil")
	}
}
