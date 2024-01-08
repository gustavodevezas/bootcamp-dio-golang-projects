// Objetivo: Identificar e imprimir os números divisíveis por 3 e 5 do 1 a 100.
// Demonstração do uso do operador de resto % e do loop for para iteração.
// Demonstração do uso do if para verificar divisibilidade.
// Utiliza o pacote fmt para formatação ao identificar e imprimir os números divisíveis por 3 e 5 dentro do laço.

// Pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Este loop percorre os números de 1 a 100.
	for i := 1; i <= 100; i++ {
		// Verifica se o número atual é divisível por 3 e 5.
		if i%3 == 0 && i%5 == 0 {
			// Se for divisível por 3 e 5, imprime PIN PAN.
			fmt.Println("PIN PAN")
		} else if i%3 == 0 {
			// Se for divisível por 3, imprime PIN.
			fmt.Println("PIN")
		} else if i%5 == 0 {
			// Se for divisível por 5, imprime PAN.
			fmt.Println("PAN")
		} else {
			// Se não for divisível por 3 ou 5, imprime o próprio número.
			fmt.Println(i)
		}
	}
}
