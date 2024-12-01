package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64
	var n int = 10000
	fmt.Println(n)

	// uint → unsyigned int (números negativs)
	var n2 uint32 = 1000
	fmt.Println(n2)

	// Alias
	// Int32 = Rune
	var n3 rune = 1234

	// Alias
	// byte = Uint8
	var n4 byte = 123
	fmt.Println(n3, n4)

	var realN float32 = 123.45
	fmt.Println(realN)

	str := "Text"
	fmt.Println(str)

	//Não tem char no go, o mais perto q temos é:
	char := 'B'
	fmt.Println(char) // B = (66), char é int

	// Valor 0, quando uma variável não é inicializada o seu valor
	var (
		texto string // "valor inicial é '' "
		nn    int16  // "valor inicial é 0"
		boobo bool   // "valor inicial é false"
	)
	fmt.Println(texto, nn, boobo)

	// Boolean
	var boolean bool = true
	boolean2 := false

	fmt.Println(boolean, boolean2)

	//Error type
	var error1 error // <nill> é praticamente null
	var erro2 error = errors.New("Error Internal")
	fmt.Println(error1, erro2)
}
