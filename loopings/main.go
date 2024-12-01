package main

import (
	"fmt"
)

func main() {
	// i := 0

	// Não tem while, isso é o mais próximo de um while
	// for i < 10 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando i:")
	// 	i++
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j: ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Davi", "Luca", "Edu"}

	// For in ou for of
	for index, value := range nomes {
		fmt.Println(index, value)
	}

	for _, value := range nomes {
		fmt.Println(value)
	}

	for index, value := range "Palavra" {
		fmt.Println(index, string(value)) //Tem q converter para string se n ele traz a posição na tabela Asp
	}

	user := map[string]string{ //map[tipo_da_chave]tipo_do_valor
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	for key, value := range user {
		fmt.Println(key, ": ", value)
	}

	//While True
	// for {
	// 	fmt.Println("Executando infitamente")
	// 	time.Sleep(time.Second)
	// }

}
