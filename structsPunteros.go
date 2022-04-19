package main

import "fmt"


type pc struct{
	ram int
	disk int
	brand string
}

func (myPc pc) ping(){
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicatedRAM(){
	myPc.ram = myPc.ram * 2

}

func main() {
	a := 50
	b := &a //la direccion donde esta guardado a (un puntero)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("--------")
	fmt.Println(b)
	fmt.Println(*b)//el & permite acceder a la direccion de memoria y el * al valor de esa direccion de memoria

	myPc := pc{ram:16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicatedRAM()

	fmt.Println(myPc)
	myPc.duplicatedRAM()

	
}