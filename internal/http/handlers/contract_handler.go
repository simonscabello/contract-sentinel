package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
	"github.com/simonscabello/contract-sentinel/internal/results"
)

type ContractHandler struct {
	Service    contracts.ContractValidator
	Repository results.Saver
}

type contractRequest struct {
	Path        string `json:"path"`
	ProviderURL string `json:"provider_url"`
	Consumer    string `json:"consumer"`
	Provider    string `json:"provider"`
	Version     string `json:"version"`
}

func (h *ContractHandler) ValidateContract(w http.ResponseWriter, r *http.Request) {
	var req contractRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Corpo inválido", http.StatusBadRequest)
		return
	}

	contract := contracts.Contract{
		Path:        req.Path,
		ProviderURL: req.ProviderURL,
		Consumer:    req.Consumer,
		Provider:    req.Provider,
		Version:     req.Version,
	}

	result, err := h.Service.ValidateContract(contract)

	_ = h.Repository.Save(context.Background(), contract, result)

	if err != nil {
		http.Error(w, result.Output, http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
