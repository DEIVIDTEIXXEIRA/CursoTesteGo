package iteracao

const quantidadeRepeticoes = 4

func Repetir(texto string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += texto
	}
	return repeticoes
}