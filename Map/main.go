package main

import "fmt"

func main() {
	fmt.Println("Map")
	user := map[string]string{ //map[tipo_da_chave]tipo_do_valor
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	user2 := map[string]map[string]string{ //map[tipo_da_chave]tipo_do_valor
		"nome": {
			"primeiro": "Arthur",
			"ultimo":   "Silva",
		},
	}

	fmt.Println(user["nome"])
	fmt.Println(user)
	fmt.Println(user2)
	// Apagar uma chave
	delete(user2, "nome")
	user2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(user2)
}
