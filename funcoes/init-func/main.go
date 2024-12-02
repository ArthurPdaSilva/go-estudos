package main

import "fmt"

var n int

func init() {
	fmt.Println("Função que é chamada antes do main")
	fmt.Println("E pode ser usado uma por arquivo, enquanto main é uma por package")
	n = 10
}

func main() {
	fmt.Println("Função main sendo usada")
	fmt.Println(n)
}
