package main

import (
	"fmt"
	"strings"
)

func converterParaMaiusculas(str string) string {
	return strings.ToUpper(str)
}

func main() {

	var str string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&str)

	resultado := converterParaMaiusculas(str)

	fmt.Println("Resultado:", resultado)
}
