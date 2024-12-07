package main

import (
	"fmt"
)

func main() {
	// canal := make(chan string)
	// canal <- "Olá mundo!"
	// //Do jeito q tá vai dar deadlock, pq as funções estão na mesma função e enviar e receber estão na mesma linha
	// mns := <-canal

	//Canal com buffer, ele só vai bloquear quando ultrapassar o limite do canal
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Programando em Go!"
	mns := <-canal
	mns2 := <-canal

	fmt.Println(mns)
	fmt.Println(mns2)

	fmt.Printf("Fim do programa!")

}
