package main

import "fmt"

//Rest operator
func soma(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

// Só pode um parâmetro variático por função e ele só pode ser a última coisa do parâmetro
func escrever(texto string, nums ...int) {
	for _, num := range nums {
		fmt.Println(texto, num)
	}
}

func main() {
	fmt.Println(soma(1, 23, 2, 34, 5, 5, 7, 46, 6))
}
