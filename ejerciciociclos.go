package main



func main(){

	counterBreak := 10

	for{
		println(counterBreak)
		if counterBreak == 0{
			break
		}
		counterBreak--
	}
}