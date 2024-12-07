package endereco

//Os testes são nativos do Golang
import "testing"

//Estrutura básica de teste em go, ter Test no início e ter o valor t *testing.T
//Só rodar o go test para rodar a suit
// Para rodar em todos os diretorios: go test ./...

func TestIsTipoValido(t *testing.T) {

	endTest := "Avenida Paulista"

	espect := true
	response := IsTipoValido(endTest)

	if espect != response {
		t.Errorf("O tipo recebido é diferente do esperado, esperado: %t e recebeu: %t", espect, response)
	}

}

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  bool
}

// Go tem caches no test
// Go test -v || Ele detalha mais no terminal
// Rodar testes em paralelo, só rodar t.Parallel() em todas as funções
//go test --cover //Mostra a cobertura dos testes
//Gerar relatório das linhas cobertas e outras não: go test --coverprofile cobertura.txt
// Ler arquivo relatório: go tool cover --func=cobertura.txt
// Ler arquivo relatório com html: go tool cover --html=cobertura.txt

func TestComVariosCenariosIsTipoValido(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", true},
		{"Avenida Paulista", true},
		{"Estrada dos Bandeirantes", true},
		{"Rodovia dos Imigrantes", true},
		{"Praça da Sé", false},
		{"Alameda Santos", false},
		{"Travessa XYZ", false},
		{"Beco Sem Saída", false},
		{"Rua", true},
		{"", false},
		{"  ", false},
		{"Rua123", false},
		{"123 Rua", false},
	}

	for _, cenario := range cenariosDeTeste {
		response := IsTipoValido(cenario.enderecoInserido)
		if cenario.retornoEsperado != response {
			t.Errorf("O tipo recebido é diferente do esperado, esperado: %t e recebeu: %t", cenario.retornoEsperado, response)
		}
	}

}

func TestQualquer(t *testing.T) {
	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}
