package dependencias

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Cris")

	resultado := buffer.String()
	esperado := "Olá, Cris"

	if resultado != esperado {
		t.Errorf("Erro: Esperado'%s', resultado '%s'.", esperado, resultado)
	}
}
