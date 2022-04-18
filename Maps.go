package main

import "fmt"

func main() {
	//Lo que en Python se conoce como diccionarios, en Go
	// son Maps

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer un Map
	for i, valor := range m{
		fmt.Println(i, valor)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}