package pr

import "fmt"

func main() {
	fmt.Println("Hello")
}

const prefixoOlaPortugues = "Olá, "

func Ola(nome, idioma string) string {
	const prefixoOlaPortugues = "Olá, "
	const espanhol = "Hola, "
	const prefixoEspanhol = "espanhol"
	const frances = "Bonjour, "
	const prefixoFrances = "frances"

	if nome == "" {
        nome = "Mundo"
    }

    prefixo := prefixoOlaPortugues

    switch idioma {
    case frances:
        prefixo = prefixoFrances
    case espanhol:
        prefixo = prefixoEspanhol
    }

    return prefixo + nome
}