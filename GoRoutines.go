package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Print(text)
}

func main() {

	var wg sync.WaitGroup //interactura de forma eficiente con GoRoutines
	//acumula y va liberando poco a poco

	fmt.Println("Hello")
	wg.Add(1)//agregamos una GoRoutine

	go say("World", &wg) //con go, se ejecuta de forma concurrente

	wg.Wait()//Que espere hasta que se ejecuten todas las GoRoutines

	//funciones anonimas

	go func (){
		fmt.Println("Adios")
	}()

	go func (text string){
		fmt.Println(text)
	}("Adios")


}