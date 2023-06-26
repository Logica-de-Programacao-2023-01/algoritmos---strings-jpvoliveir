package main

import (
	"fmt"
	"strings"
)

func contarPalavras(str string) int {
	palavras := strings.Fields(str)
	return len(palavras)
}

func main() {
	var str string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&str)

	numPalavras := contarPalavras(str)

	fmt.Println("A frase contém", numPalavras, "palavras.")
}
