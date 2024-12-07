package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	//Aqui to mandando o canal esperar um valor
	// A partir do momento q aqui recebeu um valor, só siga o projeto
	// mensagm := <-canal
	// fmt.Println(mensagm)
	// for { //Isso aqui vai dar deadlock, ele vai tá esperando dado, sendo q n tem mais dado subindo
	// 	mens, aberto := <-canal

	// 	//Tem mais canal aberto esperando dado ou canal passando dado, feche
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mens)
	// }

	//Enquanto o canal ainda tá aberto, mesma coisa do código acima
	for mns := range canal {
		fmt.Println(mns)
	}

	fmt.Printf("Fim do programa!")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//To mandando um valor para dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//Depois q terminar o looping, feche o canal
	close(canal)
}
