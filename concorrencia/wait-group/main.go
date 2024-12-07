package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Sincronizar funções go routines
	var waitGroup sync.WaitGroup

	//Temos duas funções de esperar
	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
