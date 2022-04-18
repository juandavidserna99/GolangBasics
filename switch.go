package main

import "fmt"

func main() {

	//Forma 1

	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Forma 2

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Forma 3: sin condicion

	value:= 200

	switch{
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:// es buena practica dejar un default
		fmt.Println("No hay condicion")
	}

}