package main

import "fmt"

func main() {

	// Declaracion de constantes

	//const pi float64 = 3.14
	//fmt.Println(pi)

	//const pi2 = 3.1415
	//fmt.Println(pi2)
	//fmt.Println("---------")
	//fmt.Println("Pi ", pi)
	//fmt.Println("Pi2 ", pi2)
	//fmt.Println("---------")
	
	//Declaracion de variables

	//base := 12
	//var altura int = 14
	//var area int

	//fmt.Println("---------")
	//fmt.Println(base, altura, area)
	// Zero value

	//var a int
	//var b float64
	//var c string
	//var d bool

	//fmt.Println("---------")
	//fmt.Println(a, b, c, d)

	// calcular el area de un cuadrado

	//const basecuadrado = 10
	//areacuadrado := basecuadrado*basecuadrado

	//fmt.Println("El area del cuadrado es ", areacuadrado)


	//fmt.Println("---------")
	//fmt.Println("Hello World")

	// Operadores matematicos
	x := 10
	y := 50

	//suma 
	result := x + y
	fmt.Println("La suma es", result)

	// Resta

	result = y - x
	fmt.Println("La diferencia es", result)

	//Multiplicacion

	result = x * y
	fmt.Println("La multiplicacion es", result)

	//Division

	result = y / x
	fmt.Println("La division es", result)

	//Modulo al dividir

	result = y % x
	fmt.Println("El modulo es", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental :", x)

	//Area de un rectangulo

	const n int = 5
	const m int = 10
	area1 := n*m
	fmt.Println("El area de un rectangulo es: ", area1)

	//Area de un trapecio
	const j int = 7
	area1 = (m+j)*n/2
	fmt.Println("El area de un trapecio es", area1)

	//Area de un circulo
	const pi float64 = 3.1415
	var r float64 = 3
	areacirculo := pi * r * r
	fmt.Println("El area del circulo es", areacirculo)
}