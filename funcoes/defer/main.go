package main

import "fmt"

func fun1() {
	fmt.Println("Executando a função 1")
}

func fun2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	m := (n1 + n2) / 2

	if m > 6 {
		fmt.Println("Aluno aprovado")
		return true
	}

	fmt.Println("Aluno reprovado")
	return false
}

func main() {
	defer fun1() // defer é adiar, ele vai adiar ao máximo de tempo possível
	fun2()
	alunoAprovado(6, 7)
}
