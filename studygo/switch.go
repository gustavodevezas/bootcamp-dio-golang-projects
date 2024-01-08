// Objetivo: Utilizar um loop for e switch para imprimir números de 0 a 5 com representações textuais.

// Pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Loop que percorre os números de 0 a 5.
	for i := 0; i <= 5; i++ {
		// O switch é usado para associar cada número a uma representação textual.
		switch i {
		case 0:
			// Se i for 0, imprime "Zero".
			fmt.Println(i, "(Zero)")
		case 1:
			// Se i for 1, imprime "Um".
			fmt.Println(i, "(Um)")
		case 2:
			// Se i for 2, imprime "Dois".
			fmt.Println(i, "(Dois)")
		case 3:
			// Se i for 3, imprime "Três".
			fmt.Println(i, "(Três)")
		case 4:
			// Se i for 4, imprime "Quatro".
			fmt.Println(i, "(Quatro)")
		case 5:
			// Se i for 5, imprime "Cinco".
			fmt.Println(i, "(Cinco)")
		}
	}
}
