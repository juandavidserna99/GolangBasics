package main

import "fmt"

func message(text string, c chan string){
	c <- text
}

func main() {
	c := make(chan string, 2)

	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	// Range & close

	close(c) //asi tenga capacidad no recibe otro dato mas
	//los canales se cierran cuando se dejan de usar

	for message := range c {
		fmt.Println(message)
	}
	//Cuando hay multiples canales y no se sabe cual va a responder primero
	//se usa select

	email := make(chan string)
	email2 := make(chan string)
	go message("Mensaje 1", email)
	go message("Mensaje 2", email2)
	for i:=0; i<2; i++{
		select {
		case m1:= <- email:
			fmt.Println("Email recibido de email 1", m1)
		
		case m2:= <- email2:
			fmt.Println("Email recibido de email 1", m2)
		}
	}
}