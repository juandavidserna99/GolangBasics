package main

import "fmt"

func isPalindromo(text string){
	var textReverse string

	for i := len(text) - 1; i >= 0; i--{
		textReverse += string(text[i])
	}

	if text == textReverse{
		fmt.Println("Es palindromo")
	}else{
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	//Para mostrar el indice y el valor del slice
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	fmt.Println("---------------------")

	//Para mostrar el valor
	for _ , valor := range slice {
		fmt.Println(valor)
	}

	fmt.Println("---------------------")

	//Para mostrar solo el index

	for i := range slice{
		fmt.Println(i)
	}

	//Palimdromo
	isPalindromo("ama")
}