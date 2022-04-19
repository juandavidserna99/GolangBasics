package main

import "fmt"

type papeleria struct {
	cuadernos int
	marca     string
	pago      string
}

func (miPapeleria *papeleria) cCuaderno(){

	miPapeleria.cuadernos *= 2

}

func main() {

	miPapeleria := papeleria{cuadernos: 10, marca: "Norma", pago: "Tarjeta"}

	fmt.Println(miPapeleria)
	miPapeleria.cCuaderno()

	fmt.Println(miPapeleria)
	miPapeleria.cCuaderno()
}