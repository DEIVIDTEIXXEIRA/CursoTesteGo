package pr_test

import (
	"deivid/pr"
	"testing"
)

func TestOla(t *testing.T) {
    verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
        t.Helper()
        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    }

    t.Run("diz olá para as pessoas", func(t *testing.T) {
        resultado := pr.Ola("Chris", "")
        esperado := "Olá, Chris"
        verificaMensagemCorreta(t, resultado, esperado)
    })

    t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
        resultado := pr.Ola("", "")
        esperado := "Olá, Mundo"
        verificaMensagemCorreta(t, resultado, esperado)
    })

	t.Run("em espanhol", func(t *testing.T) {
		resultado := pr.Ola("Deivid", "espanhol")
		esperado := "Hola, Deivid"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em Francês", func(t *testing.T) {
		resultado := pr.Ola("Deivid", "frances")
		esperado := "Bonjour, Deivid"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}