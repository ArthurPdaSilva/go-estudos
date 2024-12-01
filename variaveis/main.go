package main

import "fmt"

func main() {
	// Explícito
	var variavel string = "variável 1"
	fmt.Println(variavel)

	// Implícito
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Declarar mais de uma variável de uma vez
	var (
		variavel3 string = "ghmm"
		variavel4 string = "hmmmm"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel5, variavel6)

	const contExample string = "Constante 1"
	fmt.Println(contExample)

	// Trocar valores de variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
