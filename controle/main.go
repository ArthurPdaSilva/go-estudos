package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else if num == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//If init, iniciando uma variável numa condição
	if anotherNum := num; anotherNum > 0 {
		fmt.Println("Número é maior que 0")
		//anotherNum só pode ser usada no if e else
	}
}
