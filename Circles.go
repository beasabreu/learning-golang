package main

import "fmt"

const pi float64 = 3.14

var raio float64
var operacao int
var resultadoArea, resultadoPerimetro float64

func main() {
	fmt.Printf("Bem vindo a calculadora de circunferencias!\nOperacoes:\n1 - Area da circunferencia\n2 - Perimetro da circunferencia\t")
	fmt.Printf("\nDigite a operacao desejada: ")
	fmt.Scanf("%d", &operacao)
	fmt.Printf("Digite o raio da circunferencia: ")
	fmt.Scanf("%f", &raio)
	switch {
	case operacao == 1:
		resultadoArea = (raio * raio * pi)
		fmt.Printf("Seu resultado eh: %.2f\n", resultadoArea)
	case operacao == 2:
		resultadoPerimetro = (2 * pi * raio)
		fmt.Printf("Seu resultado eh: %.2f\n", resultadoPerimetro)
	default:
		fmt.Printf("Operacao invalida.")
	}
}
