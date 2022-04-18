package main

import "fmt"

//funcion que recibe un numero

func parImpar(x int)bool{
	return x % 2 == 0
} 

//funcion que recibe un string

func verifique(user, password string)bool{
	return user == "Juan David" && password == "12345"
}

func main(){

	if parImpar(2){
		fmt.Println("Es par")
	}else{
		fmt.Println("Es impar")
	}

	if verifique("Juan","12345"){
		fmt.Println("Validado")
	}else{
		fmt.Println("No valido")
	}

}