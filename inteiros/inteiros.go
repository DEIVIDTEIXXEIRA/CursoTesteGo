package inteiros

func Soma(numeros []int) int {
    soma := 0
    for _, numero := range numeros {
		soma += numero	
	}
	return soma
}

func SomaTudo(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        somas = append(somas, Soma(numeros))
    }

    return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        final := numeros[1:]
        somas = append(somas, Soma(final))
    }

    return somas
}