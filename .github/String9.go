package main

import (
	"fmt"
	"strings"
)

func substituirLetra(str, antiga, nova string) string {
	return strings.ReplaceAll(str, antiga, nova)
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	var antiga string
	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&antiga)

	var nova string
	fmt.Print("Digite a letra substituta: ")
	fmt.Scanln(&nova)

	resultado := substituirLetra(str, antiga, nova)

	// Exibir o resultado
	fmt.Println("Resultado:", resultado)
}
