package main

import (
	"log"
	"net/http"

	router "github.com/simonscabello/contract-sentinel/internal/http"
)

func main() {
	log.Println("Contract Sentinel rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", router.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
