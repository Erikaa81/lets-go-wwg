//Exercício #1
//Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.
package main

import "fmt"

func main() {
	numero := 13

	if numero == 0 {
		fmt.Println("Esse numero é zero")
	}

	if numero%2 == 0 {

		fmt.Println("Esse numero é multiplo 2")
	}

	if numero%3 == 0 {
		fmt.Println("Esse numero é multiplo 3")

	} else {
		fmt.Println("Esse numero não é 0, não e multiplo de 2 nem multiplo de 3")
	}
}
