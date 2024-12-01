package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Inferir tamanho
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Tem tamanho dinâmico
	slice := []int{11, 1, 1, 1, 23, 23, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf((slice)))
	fmt.Println(reflect.TypeOf((array1)))

	//Add novo item
	slice = append(slice, 12)
	fmt.Println(slice)

	// Pegar os itens de um array
	// índice 1 vai ser incluído e vai até o índice 2 (3 é excluído)
	slice2 := array2[1:3]
	array2[1] = "Posição alterada" // slice2 aponta para o array2, então, se eu alterar o array, ele é alterado junto
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("---------------")
	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 3)
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho retornado
	fmt.Println(cap(slice3)) // Capacidade | tamanho do array

	//Como ele viu q a capacidade ir ser estourada, ele duplica o tamanho do array interno
	fmt.Println("---------------")

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	fmt.Println(len(slice4))
	slice4 = append(slice4, 4)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	fmt.Println(len(slice4))
}
