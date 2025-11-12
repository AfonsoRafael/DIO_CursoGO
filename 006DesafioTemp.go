// declarar pacote principal
package main
// importar pacote de formatação
import "fmt"
// função principal

const ebolicao_Agua_Farenheit = 212.0
func main() {
	var temp_F float64 = ebolicao_Agua_Farenheit
	var temp_C float64 = (temp_F - 32) * 5 / 9
	fmt.Println("A temperatura de ebolição da agua °F e ", temp_F)
	fmt.Println("A temperatura de ebolição da agua °C e ", temp_C)


}