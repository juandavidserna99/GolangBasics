package main

import "fmt"

type figuras interface {
	area() float64
}

type rombo struct {
	largo float64
	corto float64
}

type cometa struct {
	largo float64
	corto float64
}

type trapecio struct {
	altura float64
	basel  float64
	basec  float64
}

func (r rombo) area() float64 {
	return r.largo * r.corto / 2
}
func (c cometa) area() float64 {
	return c.largo * c.corto / 2
}
func (t trapecio) area() float64 {
	return (t.basel + t.basec) * t.altura / 2
}

func calcular(f figuras) {
	fmt.Println("Area: ", f.area())
}

func main() {
	myRombo := rombo{largo: 6, corto: 3}
	myCometa := cometa{largo: 6, corto: 3}
	myTrapecio := trapecio{altura: 5, basel: 6, basec: 3}

	calcular(myRombo)
	calcular(myCometa)
	calcular(myTrapecio)
}