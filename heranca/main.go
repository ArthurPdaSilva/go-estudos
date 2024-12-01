package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   float32
}

type student struct {
	person  //To buscando todos os atributos de person, tipo extends
	course  string
	college string
}

func main() {
	fmt.Println("Herança")
	p := person{"João", "Pedro", 20, 1.78}
	fmt.Println(p)

	e1 := student{p, "Engenharia", "USP"}
	fmt.Println(e1)
}
