package main

import (
	"fmt"
	"strconv"
)

func verificarSequenciaDecrescente(str string) bool {
	if len(str) < 2 {
		// A sequência precisa ter pelo menos dois dígitos para ser considerada decrescente
		return false
	}

	// Converte a string em um slice de inteiros
	digitos := make([]int, len(str))
	for i, char := range str {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			// Se algum caractere não for um dígito válido, retorna falso
			return false
		}
		digitos[i] = digito
	}

	// Verifica se os dígitos formam uma sequência numérica decrescente
	for i := 1; i < len(digitos); i++ {
		if digitos[i] >= digitos[i-1] {
			return false
		}
	}

	return true
}

func main() {
	// Solicita a string ao usuário
	var str string
	fmt.Print("Digite uma string numérica: ")
	fmt.Scanln(&str)

	// Verifica se é uma sequência numérica decrescente
	if verificarSequenciaDecrescente(str) {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}
