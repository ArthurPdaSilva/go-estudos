package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Responsável por conter todas as rotas
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost) //ou Só "POST"

	fmt.Println("Conectando servidor")
	log.Fatal(http.ListenAndServe(":5001", router))
}
