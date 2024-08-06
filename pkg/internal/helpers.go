package internal

func CalculaDigitoVerificadorSP(ie []int) int {
	pesos := []int{1, 3, 4, 5, 6, 7, 8, 10}
	soma := 0
	for i := 0; i < 8; i++ {
		soma += ie[i] * pesos[i]
	}
	resto := soma % 11
	digito := 11 - resto

	if digito == 10 || digito == 11 {
		return 0
	}
	return digito
}

func CalculaDigitoVerificador(cnpj []int) int {
	pesos := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2, 6, 5, 4, 3, 2}
	if len(cnpj) == 12 {
		pesos = pesos[1:]
	}
	soma := 0
	for i := range cnpj {
		soma += cnpj[i] * pesos[i]
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}
