package main

import "fmt"

func seqFibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return seqFibo(n-2) + seqFibo(n-1)
}

//Dá para especificar, q no caso o todos só recebe dados e o results só envia dados
func worker(todos <-chan int, results chan<- int) {
	for todo := range todos {
		results <- int(seqFibo(todo))
	}
}

func main() {
	todos := make(chan int, 45)
	results := make(chan int, 45)

	// Vários processos fazendo simultâneos
	go worker(todos, results)
	go worker(todos, results)
	go worker(todos, results)
	go worker(todos, results)
	go worker(todos, results)
	go worker(todos, results)
	go worker(todos, results)
	go worker(todos, results)

	for i := 0; i < 45; i++ {
		todos <- i
	}

	close(todos)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}
