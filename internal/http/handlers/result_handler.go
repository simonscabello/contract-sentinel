package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/simonscabello/contract-sentinel/internal/results"
)

type ResultsHandler struct {
	Repository results.SaverWithQuery
}

func (h *ResultsHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	successStr := q.Get("success")
	var success *bool

	if successStr != "" {
		val, err := strconv.ParseBool(successStr)
		if err != nil {
			http.Error(w, "Parâmetro 'success' inválido", http.StatusBadRequest)
			return
		}
		success = &val
	}

	res, err := h.Repository.FindAll(context.Background(), results.QueryParams{
		Consumer: q.Get("consumer"),
		Provider: q.Get("provider"),
		Success:  success,
	})

	if err != nil {
		http.Error(w, "Erro ao buscar resultados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
