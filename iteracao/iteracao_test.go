package iteracao

import "testing"

func TestIterar(t *testing.T) {
	resultado := Repetir("a")
	esperado := "aaaa"

	if resultado != esperado {
		t.Errorf("Erro: recebido '%s'. esperado '%s'.", resultado, esperado)
		return
	}
}

func BenchmarkRepetir(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repetir("a")
    }
}



