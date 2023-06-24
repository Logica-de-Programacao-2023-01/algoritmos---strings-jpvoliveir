package main

import (
	"fmt"
	"strconv"
)

func verificarSequenciaCrescente(str string) bool {
	if len(str) < 2 {
		return false
	}

	digitos := make([]int, len(str))
	for i, char := range str {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digitos[i] = digito
	}

	for i := 1; i < len(digitos); i++ {
		if digitos[i] <= digitos[i-1] {
			return false
		}
	}

	return true
}

func main() {
	var str string
	fmt.Print("Digite uma string numérica: ")
	fmt.Scanln(&str)

	if verificarSequenciaCrescente(str) {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
