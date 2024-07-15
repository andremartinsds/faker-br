package pkg

import (
	"strconv"
	"testing"
)

func TestFaker_Numero(t *testing.T) {
	t.Run("deve criar um numero de forma randomica com size", func(t *testing.T) {
		n := G.Numero(11)
		size := len(strconv.Itoa(n))
		if 11 != size {
			t.Errorf("numero %d, want %d", n, size)
		}
	})
	t.Run("deve retornar um cep de rorma randomica", func(t *testing.T) {
		cep := G.Cep()

		if cep == "" {
			t.Errorf("cep is empty")
		}
	})
}
