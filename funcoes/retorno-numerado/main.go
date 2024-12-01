package main

import "fmt"

// Com retorno numerado, eu n preciso declarar o retorno, só dar um return no fim
func calc(n1, n2 int) (soma, subtacao int) {
	soma = n1 + n2
	subtacao = n1 - n2
	return
}

func main() {
	soma, subtracao := calc(10, 20)
	fmt.Println(soma, subtracao)
}
