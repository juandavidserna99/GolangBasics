package main

import (
	"fmt"
	"strings"
)

func isPalindromo(texto string) {
	var textoReverso string

	for i := len(texto) - 1; i >= 0; i-- {

		textoReverso += string(texto[i])
	}

	if texto == textoReverso {
		fmt.Println("Es palindromo")
	}else{
		fmt.Println("No es palindromo")
	}
}

func main() {

	isPalindromo(strings.ToLower("Ana"))
}