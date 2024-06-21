/**
Um professor de ensino médio solicitou aos seus alunos que desenvolvessem um
programa para converter o valor do ponto de ebulição da água de Kelvin para graus
Celsius.

Dica: C = K - 273
*/

package main

import "fmt"

const ebulicaoKelvin = 373.15

func main() {
	var ebulicaoCelsius = KelvinToCelsius(ebulicaoKelvin)

	fmt.Println("Ponto de ebulição em Kelvin: ", ebulicaoKelvin)
	fmt.Println("Ponto de ebulição em Celsius: ", ebulicaoCelsius)
}

func KelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
  return celsius
}