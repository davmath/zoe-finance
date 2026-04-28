package main

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao realizar leitura do arquivo .env")
	}

	database.Connect()

	http.HandleFunc("/transacoes", handlers.HandleTransacoes)
	porta := ":8000"
	fmt.Printf("API Zoe Finance started at %s port\n", porta)

	log.Fatal(http.ListenAndServe(porta, nil))
}