package main

import "fmt"

func normalFunction(message string){

	fmt.Println(message)
}

func tripleArgument(a, b int, c string){
	fmt.Println(a, b, c)
}

func retunrValue(a int) int {
	return a * 2

}

func dobleReturn(a int)(c, d int){
	return a, a*2
}

func main() {

	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := retunrValue(2)

	fmt.Println("Value: ", value)

	value1, value2 := dobleReturn(2)
	fmt.Println("Value 1 y 2:", value1, value2)
}