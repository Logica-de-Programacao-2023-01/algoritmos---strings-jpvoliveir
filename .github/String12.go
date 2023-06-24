package main

import (
	"fmt"
	"strings"
)

func verificarPalindromo(str string) bool {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")

	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	if verificarPalindromo(str) {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}
