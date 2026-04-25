package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(){
	dsn:= os.Getenv(("DATABASE_URL"))
	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil{
		log.Fatal("Erro ao configurar conexão com o banco de dados", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("O banco de dados está inacessivel: ", err)
	}

	fmt.Println("Conexão realizada com sucesso")	
}