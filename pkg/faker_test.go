package pkg

import (
	"github.com/andremartinsds/faker-br/pkg/test"
	"strconv"
	"testing"
)

func TestFaker(t *testing.T) {
	t.Run("deve criar um numero de forma randomica com size", func(t *testing.T) {
		n := G.Numero(11)
		size := len(strconv.Itoa(n))

		test.AssertLen(t, 11, size)
	})
	t.Run("deve retornar um cep de forma randomica", func(t *testing.T) {
		cep := G.Cep()

		if cep == "" {
			t.Errorf("cep is empty")
		}

		test.AssertLen(t, len(cep), 8)
	})
}
