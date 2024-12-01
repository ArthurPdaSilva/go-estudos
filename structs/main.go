package main

// Ele é uma coleção de campos
// Go não tem classes e não orientado a objetos

import "fmt"

type user struct {
	name string
	age  uint8 //números positivos apenas (classe)
	address
}

type address struct {
	location string
	number   uint8
}

func main() {
	fmt.Println("Arquivo structs")
	var u user
	fmt.Println(u) //Seta todos os campos com seus valores zerados: { 0}

	u.name = "Edu"
	u.age = 13

	fmt.Println(u)

	addressE := address{"Rua dos bobos", 2}
	u2 := user{"Davi", 21, addressE}
	fmt.Println(u2)

	u3 := user{name: "Davi"}
	fmt.Println(u3)

}
