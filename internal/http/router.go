package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	pactAdapter "github.com/simonscabello/contract-sentinel/internal/adapters/pact"
	"github.com/simonscabello/contract-sentinel/internal/contracts"
	"github.com/simonscabello/contract-sentinel/internal/http/handlers"
	"github.com/simonscabello/contract-sentinel/internal/results"
	"github.com/simonscabello/contract-sentinel/pkg/mongo"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Healthcheck
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Conexão ao Mongo e criação do repositório
	db := mongo.GetClient().Database("contract_sentinel")
	repo := results.NewRepository(db)

	// Contract validator
	adapter := pactAdapter.NewPactAdapter()
	service := contracts.NewContractValidationService(adapter)
	contractHandler := handlers.ContractHandler{
		Service:    service,
		Repository: repo,
	}

	// Results handler
	resultsHandler := handlers.ResultsHandler{
		Repository: repo,
	}

	// Rotas
	r.Post("/contracts/test", contractHandler.ValidateContract)
	r.Get("/results", resultsHandler.GetResults)

	return r
}
