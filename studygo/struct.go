// Objetivo: Demonstrar a criação e impressão de estruturas (structs).
// Define uma estrutura pessoa com os campos nome e idade.

// O pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// Define a estrutura pessoa com os campos nome do tipo string e idade do tipo inteiro.
type pessoa struct {
	nome  string
	idade int
}

// A função principal que inicia a execução do programa.
func main() {
	// Cria a instância da estrutura pessoa com valores específicos e imprime.
	fmt.Println(pessoa{"Stefany", 29})
	// Cria outra instância da estrutura pessoa com valores específicos e imprime.
	fmt.Println(pessoa{"Gustavo", 38})
	// Cria mais uma instância da estrutura pessoa com valores específicos e imprime.
	fmt.Println(pessoa{"John", 2})
}
