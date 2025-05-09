package contracts

type Contract struct {
	Path        string
	ProviderURL string
	Consumer    string
	Provider    string
	Version     string
}

type ValidationResult struct {
	Success bool
	Output  string
	Error   error
}
