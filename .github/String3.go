package main

import (
	"fmt"
	"strings"
)

func substituirCaractere(str, antigo, novo string) string {
	return strings.ReplaceAll(str, antigo, novo)
}

func main() {
	var str string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&str)

	var antigo string
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&antigo)

	var novo string
	fmt.Print("Digite o caractere substituto: ")
	fmt.Scanln(&novo)

	resultado := substituirCaractere(str, antigo, novo)

	fmt.Println("Resultado:", resultado)
}
