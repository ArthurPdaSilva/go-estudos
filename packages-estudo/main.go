package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliar.Escrever("Oi, eu sou uma função")
	error := checkmail.ValidateFormat("debnook@gmail.com")
	error2 := checkmail.ValidateFormat("123")
	fmt.Println(error)
	fmt.Println(error2)
}
