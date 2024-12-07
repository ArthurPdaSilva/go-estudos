package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrência !== Paralelismo
	go escrever("Olá mundo")
	// Esse go: independente da função acabar, vai rodando o resto do programa, não espere só va
	// go são funções que chamam métodos, mas não esperam ele terminar
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
