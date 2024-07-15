package pkg

type Faker struct {
}

func init() {
	G = &Faker{}
}

var G *Faker

func (f *Faker) Cep() string {
	return cep()
}
func (f *Faker) Cidade() string {
	return cidade()
}
func (f *Faker) Estado() string {
	return estado()
}
func (f *Faker) Bairro() string {
	return bairro()
}
func (f *Faker) UF() string {
	return uf()
}
func (f *Faker) Numero(size int) int {
	return numero(size)
}
func (f *Faker) Complemento() string {
	return complemento()
}
func (f *Faker) Observacao() string {
	return observacao()
}
