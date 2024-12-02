package main

import "fmt"

func recuperarExecucao() {
	// Recover, ele recupera a execução do problema que entra em panic
	// Panic não é do tipo do error
	if r := recover(); r != nil {
		fmt.Println("Recuperando execução")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	m := (n1 + n2) / 2

	if m > 6 {
		fmt.Println("Aluno aprovado")
		return true
	} else if m < 6 {
		fmt.Println("Aluno reprovado")
		return false

	}

	// Diferente do tratamento de error, no go se não usar recover e usar o panic, teu programa morre
	// Panic chama todos os defers daquele escopo
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}
