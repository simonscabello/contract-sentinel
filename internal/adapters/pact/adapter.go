package pact

import (
	"bytes"
	"os/exec"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
)

type PactAdapter struct{}

func NewPactAdapter() *PactAdapter {
	return &PactAdapter{}
}

func (p *PactAdapter) Validate(contractPath string, providerURL string) (contracts.ValidationResult, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("pact-verifier",
		"--provider-base-url="+providerURL,
		"--pact-url="+contractPath,
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	output := stdout.String() + stderr.String()

	return contracts.ValidationResult{
		Success: err == nil,
		Output:  output,
		Error:   err,
	}, err
}
