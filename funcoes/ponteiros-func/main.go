package main

import "fmt"

//Tamo passando a cópia de um valor
func inverterSinal(n int) int {
	return n * -1
}

// Tamo passando a referência de um valor
func inverterSinalComPonteiro(n *int) {
	*n = *n * -1
}

func main() {
	n := 20
	swapN := inverterSinal(n)
	fmt.Println(swapN)
	fmt.Println(n)

	newN := 40
	fmt.Println(newN)
	//Passei o endereço de memórica da variável
	//Quando eu quero editar o valor de newN em todo lugar
	inverterSinalComPonteiro(&newN)
	fmt.Println(newN)
}
