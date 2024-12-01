package main

import "fmt"

func main() {
	sum := 1 + 2
	subtract := 1 - 2
	division := 4 / 1
	multip := 5 * 2
	rest := 10 % 2

	fmt.Println(sum, subtract, division, multip, rest)

	var v string = "String"
	v2 := "string 2"
	fmt.Println(v, v2)

	fmt.Println(1 > 2)

	// operadores lógicas
	// fmt.Println(true && true)

	// Unários
	byn := 10
	byn++
	byn--
	byn *= 3
	byn /= 10
	byn %= 2
	fmt.Println(byn)

	// Ternários, não existe em Go
	var text string

	if byn > 5 {
		text = "Maior que 5"
	} else {
		text = "Menor que 5"
	}

	fmt.Println(text)
}
