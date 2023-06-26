package main

import (
	"fmt"
)

func encontrarLetrasUnicas(str string) string {
	letrasUnicas := ""
	frequencia := make(map[rune]int)

	for _, char := range str {
		frequencia[char]++
	}

	for _, char := range str {
		if frequencia[char] == 1 {
			letrasUnicas += string(char)
		}
	}

	return letrasUnicas
}

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	letrasUnicas := encontrarLetrasUnicas(str)

	fmt.Println("Letras únicas:", letrasUnicas)
}
