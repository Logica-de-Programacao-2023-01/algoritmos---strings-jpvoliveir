package main

import (
	"fmt"
	"unicode"
)

func contemNumero(str string) bool {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	if contemNumero(str) {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}
