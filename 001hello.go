package main  // declaração do pacote principal

import "fmt" // importação do pacote fmt para formatação de entrada e saída

func main() { // função principal onde a execução do programa começa

	fmt.Println("Hello, World!") // imprime a mensagem "Hello, World!" no console
}
/* para executar o comando via terminal faça:
   go build hello.go    isso irá gerar um arquivo executável
   ./hello              isso irá executar o arquivo gerado
   ou simplesmente use:
   go run hello.go
*/