package main

import "fmt"

var endereco string
var telefone = "9999-9999"
var quantidade int
var comprou bool // booleano! 1 ou 0, true ou false!
var valor float64
var palavras rune //runas: caracteres unicode

/* declarando variáveis: var + nome da variável + tipo! */
func main() {
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("Endereco: %s\r\n", endereco)
	fmt.Printf("Telefone: %s\r\n", telefone)
	fmt.Printf("Comprou? %v\r\n", comprou)
	fmt.Printf("Valor: %f\r\n", valor)
	fmt.Printf("Palavras: %v\r\n", palavras)

}

/* se eu imprimir uma variável sem atribuir um valor, a golang já atribui um valor " " a ela.
por exemplo, neste exercício não atribui valor algum para minhas variáveis. go irá atribuir:
endereco = " "
quantidade = 0
comprou = false
valor = 0.00
palavras = 0 */
