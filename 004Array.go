package main

import "fmt"

func main(){
var x [5]int          // Declaração de um array de inteiros com tamanho 5
x[0] = 10             // Atribuição de valores aos elementos do array
x[1] = 20
x[2] = 30
x[3] = 40
x[4] = 50
fmt.Println(x)      // Impressão do array completo
fmt.Println("Elemento na posição 2:", x[2]) // Acesso a um elemento específico do array
fmt.Println("Tamanho do array:", len(x)) // Obtém o tamanho do array

var x3 [10]int // Declaração de um array de inteiros com tamanho 10
fmt.Println(x3)// Impressão do array completo (inicializado com zeros)

var y = [3]string{"Go", "is", "fun"} // Declaração e inicialização de um array de strings
fmt.Println(y) // Impressão do array de strings

// Iteração sobre os elementos do array usando for
for i := 0; i < len(x); i++ {
	fmt.Printf("Elemento na posição %d: %d\n", i, x[i])

}


}