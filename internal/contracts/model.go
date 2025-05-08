package contracts

type Contract struct {
	Path        string
	ProviderURL string
}

type ValidationResult struct {
	Success bool
	Output  string
	Error   error
}
