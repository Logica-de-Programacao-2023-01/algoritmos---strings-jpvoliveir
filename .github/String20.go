package main

import (
	"fmt"
	"unicode"
)

func verificarCamelCase(str string) (bool, int) {
	palavras := 1
	camelCase := true

	if unicode.IsUpper(rune(str[0])) {
		camelCase = false
	}

	for i := 1; i < len(str); i++ {
		if unicode.IsUpper(rune(str[i])) {
			palavras++
		}
	}

	return camelCase, palavras
}

func main() {

	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	camelCase, palavras := verificarCamelCase(str)

	if camelCase {
		fmt.Println("A string está em camelCase.")
	} else {
		fmt.Println("A string não está em camelCase.")
	}

	fmt.Println("Número de palavras:", palavras)
}
