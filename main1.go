package main

import "fmt"

func main() {

	//Declaracion de variables fmt
	helloMessage := "Hello"
	worldMessage := "World"

	//Println
	fmt.Println(helloMessage,worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}