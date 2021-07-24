//Exercício #01
//Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.package main
package main

import "fmt"

func main() {
	a := 3000
	b := 3021
	c := 3100

	switch {
	case a > b && a > c:

		fmt.Printf("A variavel 'a' tem o mairo valor\n")
		return
	case b > c && b > a:

		fmt.Printf("A variavel 'b' tem o mairo valor\n")
		return

	default:
		fmt.Printf("A variavel 'c' tem o mairo valor\n")
	}
}
