package main

import (
	"fmt"
	"strings"
)

func verificarSubstring(str, substr string) bool {
	return strings.Contains(str, substr)
}

func main() {
	var str, substr string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&substr)

	if verificarSubstring(str, substr) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}
