package main

import "fmt"

func seqFibo(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}
	return seqFibo(n-2) + seqFibo(n-1)
}

func main() {
	fmt.Println("Função recursiva")

	for i := uint(0); i < 12; i++ {
		fmt.Print(seqFibo(i), "-")
	}
}
