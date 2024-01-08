// Objetivo: Identificar e imprimir se os números de 1 a 10 são pares ou ímpares.
// Utilização das estruturas condicionais if e else para tomada de decisões.

// O pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Loop que percorre os números de 1 a 10.
	for i := 1; i <= 10; i++ {
		// Verifica se o número atual é par ou ímpar usando o operador de resto %.
		if i%2 == 0 {
			// Se for par, imprime que é par.
			fmt.Println(i, "É Par")
		} else {
			// Se for ímpar, imprime que é ímpar.
			fmt.Println(i, "É Ímpar")
		}
	}
}
