package main

import (
	"fmt"
	"strings"
)

func removerEspacos(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func main() {
	var str string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&str)

	resultado := removerEspacos(str)

	fmt.Println("Resultado:", resultado)
}
