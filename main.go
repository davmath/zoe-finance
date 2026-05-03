package main

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/handlers"
	"fmt"
	"log"
	"net/http"

	_ "davmath/zoe-finance/docs"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Zoe Finance API
// @version 1.0
// @description API para gerenciamento de finanças.
// @host localhost:8000
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao realizar leitura do arquivo .env")
	}

	database.Connect()
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/transacoes", handlers.HandleTransacoes)
	http.HandleFunc("/responsaveis", handlers.HandleResponsavelConta)
	http.HandleFunc("/compras-parceladas", handlers.HandleComprasParceladas)
	http.HandleFunc("/cartoes-credito", handlers.HandleCartoesCredito)

	porta := ":8000"
	fmt.Printf("API Zoe Finance started at %s port\n", porta)

	log.Fatal(http.ListenAndServe(porta, nil))
}