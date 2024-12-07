package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Usuario struct {
	Id       int      `json:"id"` //Assim ele reconhece e faz a conversão
	Nome     string   `json:"nome"`
	Email    string   `json:"email"`
	Idade    int      `json:"idade"`
	Endereco Endereco `json:"endereco"`
}

type Endereco struct {
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Cep    string `json:"cep"`
}

type UserC2 struct {
	Id   int    `json:"id"` //Se deixar - ao invés de id, ele ignora esse campo na conversão
	Nome string `json:"nome"`
}

func main() {
	c := Usuario{
		Id:    1,
		Nome:  "Edu",
		Email: "edu@example.com",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua A",
			Numero: 123,
			Cidade: "Cidade B",
			Estado: "Estado C",
			Cep:    "12345-678",
		},
	}

	//marshal transforma em json
	//Array de bytes
	UserToJSON(c)

	//Map para json:
	user2 := map[string]interface{}{
		"id":   5,
		"nome": "Gimenes",
	}

	UserToJSON(user2)

	// Inverso
	cInJSON := `{"id": 5, "nome": "Rex"}`
	var user UserC2

	if err := json.Unmarshal([]byte(cInJSON), &user); err != nil {
		log.Fatal("Error")
	}

	fmt.Println(user)

}

// Converte user para json
func UserToJSON(c interface{}) {
	userInJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal("Error")
	}

	fmt.Println(userInJSON)
	fmt.Println(bytes.NewBuffer(userInJSON))
}
