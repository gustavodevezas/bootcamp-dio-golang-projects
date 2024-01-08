// Objetivo: Identificar e imprimir os números divisíveis por 3 do 1 a 100.
// Demonstração do uso do operador de resto % e do loop for para iteração.
// Demonstração do uso do if para verificar divisibilidade.
// Utiliza o pacote fmt para formatação ao identificar e imprimir os números divisíveis por 3 dentro do laço.

// Pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Este loop percorre os números de 1 a 100.
	for i := 1; i <= 100; i++ {
		// Verifica se o número atual é divisível por 3.
		if i%3 == 0 {
			// Se for divisível por 3, imprime o número.
			fmt.Println(i)
		}
	}
}
