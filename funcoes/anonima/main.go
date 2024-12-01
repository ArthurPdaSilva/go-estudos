package main

import "fmt"

func main() {
	//Função anônima
	func() {
		fmt.Println("Hello World")
	}()
	func(text string) {
		fmt.Printf("Recebido -> %s", text)
	}("Passando")

	retorno := func(text string) string {
		return text
	}("Passando")

	fmt.Printf("Variável: %s", retorno)

}
