package main

import "fmt"

//Tudo atende essa interface, por isso é genérica
func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(1)
	generic(true)

	// O any do golang é o interface{}

	// Exemplo de bagunça
	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "string",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
