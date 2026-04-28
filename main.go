package main

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao realizar leitura do arquivo .env")
	}

	fmt.Println("API Zoe Finance started!")

	database.Connect()

	fmt.Println("Testando busca de transações")

	valorMinimo := 50.0
	valorMaximo := 250.0
	
	filtro := models.FiltroTransacao{
		ValorMin: &valorMinimo,
		ValorMax: &valorMaximo,
	}

	transacoes, err := repository.BuscarTransacoes(filtro)
	if err != nil{
		log.Fatal("Erro ao buscar transações no banco de dados: ", err)
	}

	fmt.Printf("A busca retornou %d transações:\n\n", len(transacoes))

	for _, t := range transacoes{
		fmt.Printf("[ID: %d] %s | R$ %.2f | Categoria: %d\n", t.ID, t.Descricao, t.Valor, t.IDCategoria)
	}

}