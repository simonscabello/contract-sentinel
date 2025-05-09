package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	router "github.com/simonscabello/contract-sentinel/internal/http"
	mongodb "github.com/simonscabello/contract-sentinel/pkg/mongo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não carregado, usando variáveis do ambiente")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI não definida")
	}

	mongodb.Connect(mongoURI)

	log.Println("Contract Sentinel rodando na porta 8040...")
	if err := http.ListenAndServe(":8040", router.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
