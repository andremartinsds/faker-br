# faker-br

## How to install
`go get github.com/andremartinsds/faker-br`

## How to use

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
