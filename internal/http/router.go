package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"

	pactAdapter "github.com/simonscabello/contract-sentinel/internal/adapters/pact"
	"github.com/simonscabello/contract-sentinel/internal/contracts"
	"github.com/simonscabello/contract-sentinel/internal/http/handlers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Healthcheck
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Contract validator
	adapter := pactAdapter.NewPactAdapter()
	service := contracts.NewContractValidationService(adapter)
	handler := handlers.ContractHandler{Service: service}

	r.Post("/contracts/test", handler.ValidateContract)

	return r
}
