package main

import "fmt"

func main() {

	// Declaracion de constantes

	const pi float64 = 3.14
	fmt.Println(pi)

	const pi2 = 3.1415
	fmt.Println(pi2)
	fmt.Println("---------")
	fmt.Println("Pi ", pi)
	fmt.Println("Pi2 ", pi2)
	fmt.Println("---------")
	
	//Declaracion de variables

	base := 12
	var altura int = 14
	var area int

	fmt.Println("---------")
	fmt.Println(base, altura, area)
	// Zero value

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("---------")
	fmt.Println(a, b, c, d)

	// calcular el area de un cuadrado

	const basecuadrado = 10
	areacuadrado := basecuadrado*basecuadrado

	fmt.Println("El area del cuadrado es ", areacuadrado)


	fmt.Println("---------")
	fmt.Println("Hello World")
}