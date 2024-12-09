package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Eu to importando aq para o database/sql utilizar o driver
	//No caso eu importei aqui para a database/sql usar
)

// go get lib_name

func main() {
	//Executar o build: .\crud.exe
	connectionString := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connectionString)

	//Erros muitos inesperados, não necessariamente é relacionado a conexão efetivamente conectou ou não
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Print("Dentro do ping")
		log.Fatal(err)
	}

	fmt.Println("Conexão está aberta!")

	linhas, err := db.Query("select * from usuarios;")

	if err != nil {
		fmt.Println("Error na query")
		log.Fatal(err)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
