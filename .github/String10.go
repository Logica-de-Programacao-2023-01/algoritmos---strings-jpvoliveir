package main

import (
	"fmt"
	"sort"
	"strings"
)

func verificarAnagrama(str1, str2 string) bool {
	// Remover espaços em branco e converter para letras minúsculas
	str1 = strings.ReplaceAll(strings.ToLower(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToLower(str2), " ", "")

	// Verificar se as strings têm o mesmo comprimento
	if len(str1) != len(str2) {
		return false
	}

	// Converter as strings em slices de runes
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	// Ordenar as slices de runes
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	// Verificar se as strings ordenadas são iguais
	return string(runes1) == string(runes2)
}

func main() {
	// Receber as strings do usuário
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Verificar se as strings são anagramas
	if verificarAnagrama(str1, str2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
