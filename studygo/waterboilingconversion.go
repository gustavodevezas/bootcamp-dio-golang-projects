// Objetivo: Converter a temperatura do ponto de ebulição da água de Kelvin para Graus Celsius.

// O pacote main é essencial para a execução do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import "fmt" // Importa o pacote fmt para formatação de entrada e saída.

// Constante que armazena a temperatura de ebulição da água em Kelvin.
const ebulicaoAguaK float64 = 373.15

// A função principal que inicia a execução do programa.
func main() {
	tempK := ebulicaoAguaK

	// Converte a temperatura de Kelvin para Graus Celsius.
	tempC := (tempK - 273.15)

	// Imprime a temperatura convertida no console.
	fmt.Println("O valor do ponto de ebulição da Água convertendo de Kelvin para Graus Celsius é de:", tempC, "ºC")
}
