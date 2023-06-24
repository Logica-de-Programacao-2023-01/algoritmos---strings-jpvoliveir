package main

import (
	"fmt"
	"strings"
)

func substituirVogaisPorAsteriscos(str string) string {
	vogais := "aeiouAEIOU"
	for _, v := range vogais {
		str = strings.ReplaceAll(str, string(v), "*")
	}
	return str
}

func main() {
	// Solicita a string ao usuário
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Substitui as vogais por asteriscos
	resultado := substituirVogaisPorAsteriscos(str)

	// Exibe o resultado
	fmt.Println("Resultado:", resultado)
}
