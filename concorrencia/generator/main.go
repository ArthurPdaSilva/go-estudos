package main

import (
	"fmt"
	"time"
)

//Generator você encapsular uma goroutine

func main() {
	c := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func escrever(text string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return c
}
