package main

import (
	"fmt"
)

func inverterString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Receber a string do usuário
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Inverter a ordem dos caracteres
	resultado := inverterString(str)

	// Exibir o resultado
	fmt.Println("Resultado:", resultado)
}
