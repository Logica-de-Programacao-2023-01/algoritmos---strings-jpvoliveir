package main

import (
	"fmt"
	"strings"
)

func removerVogais(str string) string {
	vogais := "aeiouAEIOU"
	for _, v := range vogais {
		str = strings.ReplaceAll(str, string(v), "")
	}
	return str
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	resultado := removerVogais(str)

	fmt.Println("Resultado:", resultado)
}
