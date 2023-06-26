package main

import (
	"fmt"
)

func verificarContemSomenteNumeros(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func main() {

	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	if verificarContemSomenteNumeros(str) {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}
