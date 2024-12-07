package endereco

import "strings"

// Função que verifica se o endereço é válido
func IsTipoValido(endereco string) bool {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalvara := strings.Split(strings.ToLower(endereco), " ")[0]
	return contains(primeiraPalvara, tiposValidos)
}

//Verifica se valor existe no array
func contains(endereco string, tiposValidos []string) bool {

	for _, tipo := range tiposValidos {
		if tipo == endereco {
			return true
		}
	}
	return false
}
