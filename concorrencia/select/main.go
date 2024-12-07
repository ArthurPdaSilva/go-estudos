package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			c1 <- "c1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "c2"
		}
	}()

	for {
		//Quando uma função demora muito mais tempo do que uma para rodar, ele vê se já chegou no canal e printa, sem esperar a outra
		select {
		case m1 := <-c1:
			fmt.Printf(m1)
		case m2 := <-c2:
			fmt.Printf(m2)
		}

	}
}
