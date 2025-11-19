package main

import "fmt"
func main() {
	var x [5]int
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	for i:=0; i<len(x); i++ {
		fmt.Printf("Elemento na posição %d: %d\n", i, x[i])
	}

	//Fatias é um pedaço de um array
	fatia:= x[1:4] // Pega do índice 1 ao 3 (4-1)
	fmt.Println("Fatia do array:", fatia)

	fatias := make([]int, 5)
	fmt.Println((fatias))


}