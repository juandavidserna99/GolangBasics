package main

import "fmt"

func main() {

	//Array
	var array [4]int
	fmt.Println(array)
	//Las arrays en Go son inmutables, pero si editables:

	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	//len dice la cantidad de elementos 
	//cap dice la capacidad de guardar elementos
	fmt.Println(array, len(array), cap(array))

	//Slices
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Los slices y los arrays cumples y hacen lo mismo, solo que
	//un array es fijo e inveriable, un slice es lo contrario.
	fmt.Println("-----------------")
	//Metodos en el slice:
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Como agregar elementos en un slice con Append
	fmt.Println("-----------------")
	slice = append(slice, 7)
	fmt.Println(slice)

	//Agregar una lista

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}