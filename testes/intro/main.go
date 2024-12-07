package main

import (
	"fmt"
	"intro-testes/intro/endereco"
)

func main() {
	address := "Avenida Paulista"
	if !endereco.IsTipoValido(address) {
		fmt.Println("Tipo inválido")
		return
	}

	fmt.Print(address)
}
