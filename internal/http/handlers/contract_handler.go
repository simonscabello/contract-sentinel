package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
)

type ContractHandler struct {
	Service contracts.ContractValidator
}

type contractRequest struct {
	Path        string `json:"path"`
	ProviderURL string `json:"provider_url"`
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
	}

	result, err := h.Service.ValidateContract(contract)
	if err != nil {
		http.Error(w, result.Output, http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
