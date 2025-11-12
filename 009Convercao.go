// declarar pacote principal
package main
// importar pacote de formatação
import "fmt"
// função principal

var a int = 10
var b string = "Hello"
func main() {
var c float64 = float64(a) // conversão de int para float64
var d int = int(c) // conversão de float64 para int
fmt.Printf("Valor de a (int):%g tipo variavel %T\n", a,a)
fmt.Printf("Valor de b (string):%s tipo variavel %T\n", b,b)
fmt.Printf("Valor de c (float64) após conversão de a:%F tipo variavel %T\n", c,c)
fmt.Printf("Valor de d (int) após conversão de c:%f tipo variavel %T\n", d,d)

}