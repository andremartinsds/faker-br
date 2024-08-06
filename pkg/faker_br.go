package pkg

import (
	"fmt"
	"github.com/andremartinsds/faker-br/pkg/internal"
	"math/rand"
	"time"
)

type Faker struct {
}

func init() {
	G = &Faker{}
}

// Variavel Global para metodos fakers
var G *Faker

// Retorna um cep aleatório
func (f *Faker) Cep() string {
	return internal.CepRange()
}

// Retorna uma cidade aleatória
func (f *Faker) Cidade() string {
	return internal.CidadeRange()
}

// Retorna um estado aleatório
func (f *Faker) Estado() string {
	return internal.EstadoRange()
}

// Retorna um bairro aleatório
func (f *Faker) Bairro() string {
	return internal.BairroRange()
}

// Retorna uma unidade federativa aleatória
func (f *Faker) UF() string {
	return internal.UfRange()
}

// Retorna um numero a partir de um size adicionado como parâmetro
func (f *Faker) Numero(size int) int {
	return internal.NumeroPorTamanho(size)
}

// Retorna um complemento aleatório
func (f *Faker) Complemento() string {
	return internal.ComplementoRange()
}

// Retorna uma observação aleatória
func (f *Faker) Observacao() string {
	return internal.ObservacaoRange()
}

// Retorna um cnpj aleatório
func (f *Faker) Cnpj() string {
	rand.Seed(time.Now().UnixNano())

	var cnpj [14]int
	for i := 0; i < 12; i++ {
		cnpj[i] = rand.Intn(10)
	}

	cnpj[12] = internal.CalculaDigitoVerificador(cnpj[:12])

	cnpj[13] = internal.CalculaDigitoVerificador(cnpj[:13])

	return fmt.Sprintf("%02d.%03d.%03d/%04d-%02d",
		cnpj[0], cnpj[1]*100+cnpj[2]*10+cnpj[3],
		cnpj[4]*100+cnpj[5]*10+cnpj[6], cnpj[7]*1000+cnpj[8]*100+cnpj[9]*10+cnpj[10],
		cnpj[12]*10+cnpj[13])
}

// Retorna inscrição estadual aleatória de SP
func (f *Faker) GeraInscricaoEstadualSP() string {
	rand.Seed(time.Now().UnixNano())

	var ie [12]int
	for i := 0; i < 11; i++ {
		ie[i] = rand.Intn(10)
	}

	ie[11] = internal.CalculaDigitoVerificadorSP(ie[:11])

	return fmt.Sprintf("%03d.%03d.%03d.%03d",
		ie[0]*100+ie[1]*10+ie[2], ie[3]*100+ie[4]*10+ie[5],
		ie[6]*100+ie[7]*10+ie[8], ie[9]*100+ie[10]*10+ie[11])
}
