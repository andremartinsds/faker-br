# Faker-br ![Bandeira do Brasil](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Flag_of_Brazil.svg/50px-Flag_of_Brazil.svg.png)

<p> Pequena lib para ajudar a adicionar dados aleatórios </p>

## Como instalar
`go get github.com/andremartinsds/faker-br`

## Como usar

#### Variavel Global para metodos fakers
`var G *Faker`

#### Retorna um cep aleatório
`G.Cep()` 

#### Retorna uma cidade aleatória
`G.Cidade()` 

#### Retorna um estado aleatório
`G.Estado()` 

#### Retorna um bairro aleatório
`G.Bairro()` 

#### Retorna uma unidade federativa aleatória
`G.UF()` 

#### Retorna um numero a partir de um size adicionado como parâmetro
`G.Numero(size int)` 

#### Retorna um complemento aleatório
`G.Complemento()`

#### Retorna uma observação aleatória
`G.Observacao()`

#### Retorna um cnpj aleatório
`G.Cnpj()`

#### Retorna inscrição estadual aleatória de SP
`G.GeraInscricaoEstadualSP()`


## Modo detalhado
```golang

package anypackage
import (
	faker "github.com/andremartinsds/faker-br/pkg"
	"strconv"
)

type Endereco struct {
	Id              string
	Cep             string
	Estado          string
	Cidade          string
	Endereco        string
	Numero          int
	Complemento     string
	Bairro          string
	PontoReferencia string
	Observacao      string
}

func EnderecoDummy() *Endereco {
	return &Endereco{
		Id:             strconv.Atoi(faker.G.Numero(36)),
		Cep:             faker.G.Cep(),
		Estado:          faker.G.Estado(),
		Cidade:          faker.G.Cidade(),
		Endereco:        "other description prop",
		Numero:          faker.G.Numero(2),
		Complemento:     faker.G.Complemento(),
		Bairro:          faker.G.Bairro(),
		PontoReferencia: "any valid address",
		Observacao:      faker.G.Observacao(),
	}
}
```
