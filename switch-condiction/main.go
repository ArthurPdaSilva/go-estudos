package main

import "fmt"

func diaDaSemana(n int) string {
	switch n {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(n int) string {
	var diaDaSemana string

	// O break já inserido dentro do switch automaticamente.
	switch {
	case n == 1:
		diaDaSemana = "Domingo"
		fallthrough // Joga o código para dentro da próxima condição
	case n == 2:
		diaDaSemana = "Segunda"
	case n == 3:
		diaDaSemana = "Terça"
	case n == 4:
		diaDaSemana = "Quarta"
	case n == 5:
		diaDaSemana = "Quinta"
	case n == 6:
		diaDaSemana = "Sexta"
	case n == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia = diaDaSemana2(1)
	fmt.Println(dia)
}
