// declarar pacote principal
package main
// importar pacote de formatação
import "fmt"
// função principal

const ebolicao_Agua_Farenheit = 212.0
func main() {
	var temp_F float64 = ebolicao_Agua_Farenheit
	var temp_C float64 = (temp_F - 32) * 5 / 9
	fmt.Printf("A temperatura de ebolição da agua °F e %.2f\nA temperatura de ebolicao da agua em Cº e %g", temp_F, temp_C)

	// pode usar Println ou Printf

	// pode ser usado \n para pular linha no Printf

	// %g é um especificador de formato usado em Go para representar números de ponto flutuante de forma compacta.

	// %.2f é um especificador de formato usado em Go para representar números de ponto flutuante com duas casas decimais.


}