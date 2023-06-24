package main

import (
	"fmt"
)

func inverterString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	strInvertida := inverterString(str)

	fmt.Println("String invertida:", strInvertida)
}
