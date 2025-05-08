package contracts

type ValidatorAdapter interface {
	Validate(contractPath string, providerURL string) (ValidationResult, error)
}

type ContractValidator interface {
	ValidateContract(Contract) (ValidationResult, error)
}

type ContractValidationService struct {
	adapter ValidatorAdapter
}

func NewContractValidationService(adapter ValidatorAdapter) *ContractValidationService {
	return &ContractValidationService{adapter: adapter}
}

func (s *ContractValidationService) ValidateContract(c Contract) (ValidationResult, error) {
	return s.adapter.Validate(c.Path, c.ProviderURL)
}
