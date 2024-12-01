package main

import "fmt"

// vai somar dois valores
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Funções podem ter mais de um retorno no Go
func calcM(n1, n2 int8) (int8, int8) { //se tiver só o último parâmetro, os outros  parâmetros pegam o valor
	sum := n1 + n2
	subtract := n1 - n2

	return sum, subtract
}

func main() {
	sum := somar(10, 20)
	fmt.Println(sum)

	var f = func() {
		fmt.Println("Função F")
	}
	var f2 = func(txt string) {
		fmt.Println(txt)
	}

	f()
	f2("Resultado")

	resultsSum, resulstsSub := calcM(10, 15)
	results, _ := calcM(10, 15) // O segundo resultado é ignorado ao usar "_"

	fmt.Println(resultsSum, resulstsSub)
	fmt.Println(results)
}
