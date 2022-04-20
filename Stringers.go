package main

import "fmt"

type pc struct {
	brand string
	disk  int
	ram   int
}

func (myPC pc) String()string{
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {

	myPC := pc{brand: "msi", disk: 100, ram: 16}

	fmt.Println(myPC)
}