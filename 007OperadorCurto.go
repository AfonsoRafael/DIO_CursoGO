package main

import "fmt"

// Operador curto de declaração e atribuição := é usado para declarar e inicializar variáveis em uma única linha. 
// Ele é uma forma concisa de criar variáveis sem a necessidade de usar a palavra-chave var explicitamente.
// So  pode ser usado dentro de funções, não no escopo do pacote.
const ebolicao_Agua_Farenheit = 212.0  // Fora do main, não pode usar :=
func main(){

	temp_F := ebolicao_Agua_Farenheit
	temp_C := (temp_F - 32) * 5 / 9
	fmt.Println("A temperatura de ebolição da agua °F e ", temp_F)
	fmt.Println("A temperatura de ebolição da agua °C e ", temp_C)


}