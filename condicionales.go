package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	}else{
		fmt.Println("No es 1")
	}

	//AND &&

	if valor1 == 1 && valor2== 2{
		fmt.Println("Es verdad")
	}

	//OR ||

	if valor1 == 0 || valor2 == 2{
		fmt.Println("Es verdad en OR")
	}

	// Convertir texto a numero
	//strconv.Atoi

	value, err := strconv.Atoi("53")
	if err != nil{//nil es la forma como Go identifica un error
		log.Fatal(err) //log.Fatal indica que se cierra el codigo al no cumplir con la condicion

	}
	fmt.Println("Value: ", value)
}