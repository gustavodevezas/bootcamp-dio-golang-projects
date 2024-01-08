// Objetivo: Demonstrar a criação e utilização de fatias (slices).

// Pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Cria um array com 5 elementos do tipo float64.
	arr := [5]float64{1, 2, 3, 4, 5}

	// Cria uma fatia (slice) que inclui todos os elementos do array.
	fatia := arr[0:5]

	// Imprime a fatia, que contém todos os elementos do array.
	fmt.Println(fatia)
}
