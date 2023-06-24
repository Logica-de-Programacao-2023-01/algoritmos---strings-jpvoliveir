package main

import (
	"fmt"
	"strconv"
)

func verificarNumeroFlutuante(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func main() {
	var str string
	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&str)

	if verificarNumeroFlutuante(str) {
		fmt.Println("É um número em ponto flutuante válido.")
	} else {
		fmt.Println("Não é um número em ponto flutuante válido.")
	}
}
