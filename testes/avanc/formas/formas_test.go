package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//Ele testará as duas
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida é diferente da área esperada")
			//t.FatalF() se falhou essa linha, ele para de executar os outros testes
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida é diferente da área esperada")
			//t.FatalF() se falhou essa linha, ele para de executar os outros testes
		}
	})
}
