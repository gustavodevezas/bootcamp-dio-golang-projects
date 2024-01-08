// Objetivo: Simular a passagem de um dia (24 horas), exibindo as horas, minutos e segundos.

// Pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// A função principal que inicia a execução do programa.
func main() {
	// Variável para rastrear o total de horas passadas.
	horasPassadas := 0

	// Loop externo representa todas as horas do dia.
	for horas := 0; horas < 24; horas++ {
		fmt.Println("Horas: ", horas)
		// Loop interno representa os minutos em cada hora.
		for min := 0; min < 60; min++ {
			fmt.Println("Minutos: ", min)
		}
		// Loop interno representa os segundos em cada minuto.
		for seg := 0; seg < 60; seg++ {
			fmt.Println("Segundos: ", seg)
		}
		// Incrementa o total de horas passadas.
		horasPassadas++
	}
	// Imprime uma mensagem indicando que um dia se passou e mostra o total de horas percorridas.
	fmt.Println("Um dia se passou. Horas Totais Percorridas:", horasPassadas, "h")
}
