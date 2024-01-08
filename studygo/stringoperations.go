// Objetivo: Executar operações básicas com strings.
// Imprimir o comprimento de uma string e o byte em uma posição específica da string.
// Realiza a concatenação de duas strings.

// Pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Imprime o comprimento da string "Olá Mundo".
	fmt.Println(len("Olá Mundo"))

	// Imprime o byte na posição 2 da string "Olá Mundo".
	fmt.Println("Olá Mundo"[2])

	// Concatena as strings "Olá" e "Mundo" e imprime o resultado.
	fmt.Println("Olá" + "Mundo")
}
