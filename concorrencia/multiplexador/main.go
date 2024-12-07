package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Pega dois canais e junta em um só
	c := multi(escrever("Olá mundo!"), escrever("Olá mundo 2!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func multi(c1, c2 <-chan string) <-chan string {
	cSaida := make(chan string)

	go func() {
		for {
			select {
			case mns := <-c1:
				cSaida <- mns
			case mns := <-c2:
				cSaida <- mns
			}
		}
	}()

	return cSaida

}

func escrever(text string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return c
}
