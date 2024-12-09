package main

import "fmt"

func main() {
	// & (endereço de): Obtem o endereço de memória de uma variável.
	// * (desreferência): Acessa o valor armazenado no endereço de memória.
	fmt.Println("Ponteiros")
	v := 10
	v2 := v
	fmt.Println(v, v2)

	v++
	fmt.Println(v, v2)

	// Ponteiro é uma referência de memória
	var v3 int
	var ponteiro *int // Ele guarda o endereço de memória um tipo int

	v3 = 100
	ponteiro = &v3 // Salvando a posição de uma variável

	fmt.Println(v3, ponteiro)
	// Para pegar o valor que tá naquele endereço de memória
	fmt.Println(v3, *ponteiro)

	v3 = 150
	fmt.Println(v3, ponteiro)
	fmt.Println(v3, *ponteiro)
}
