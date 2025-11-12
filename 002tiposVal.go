/*
	   Tipos de dados em Go:
	   uint8, uint16, uint32, uint64 - inteiros sem sinal de 8, 16, 32 e 64 bits  SO NUMERO POSITIVO
	   int8, int16, int32, int64       - inteiros com sinal de 8, 16, 32 e 64 bits		NUMERO POSITIVO E NEGATIVO
	   byte                             - alias para uint8
	   rune                             - alias para int32, representa um ponto de código Unicode
	   bool                             - tipo booleano (true ou false)
	   string                           - sequência de caracteres
	   float32, float64                - números de ponto flutuante de 32 e 64 bits
*/
package main

import "fmt"

func main() {
	                //Numeros inteiros e operações básicas
	fmt.Println("1+1 =", 1+1) // Operação de adição
	fmt.Println("5-2 =", 5-2) // Operação de subtração
	fmt.Println("3*4 =", 3*4) // Operação de multiplicação
	fmt.Println("8/2 =", 8/2) // Operação de divisão
	fmt.Println("7%3 =", 7%3) // Operação de módulo (resto da divisão)
	
			                //Strings
	str1 := "Hello, "
	str2 := "Go!"
	concatenated := str1 + str2 // Concatenação de strings
	fmt.Println("Concatenação String:", concatenated)
	fmt.Println("tamanho da concatenação string:", len(concatenated)) // Comprimento da string
	fmt.Println(concatenated[2]) // Acessa o terceiro caractere da string (índice começa em 0), retornando o valor em byte (ASCII), neste caso 'l' = 108

	                      //Booleanos
	bool1 := true
	bool2 := false
	fmt.Println("AND lógico:", bool1 && bool2) // Operação AND lógico
	fmt.Println("OR lógico:", bool1 || bool2)  // Operação OR lógico
	fmt.Println("NOT lógico:", !bool1)         // Operação NOT lógico

	                      //Inferência de tipos, o Go determina o tipo automaticamente

	variavelInferred := 100 // O Go infere que é do tipo int
	fmt.Println("Variável com tipo inferido:", variavelInferred)
	fmt.Printf("Tipo da variávelInferred: %T\n", variavelInferred) // Exibe o tipo da variável
}

