package main

import "fmt"

// Area de un rectangulo

func areaRectangulo(base, altura float64 ) float64{
	return base * altura
}

// Area de un trapecio

func areaTrapecio(base1, base2, altura float64) float64{
	return (base1 + base2) * altura/2
}

// Area de un circulo

func areaCirculo(radio float64)float64{
	return 3.1415 * radio * radio
}

func main() {

	
	fmt.Println("El area del rectangulo es: ", areaRectangulo(5,2))
	fmt.Println("--------------------")

	fmt.Println("El area del trapecio es: ", areaTrapecio(5,3,2))
	fmt.Println("--------------------")

	fmt.Println("El area de un circulo es: ", areaCirculo(3))

}