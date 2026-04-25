package main

import (
	"davmath/zoe-finance/database"
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
}