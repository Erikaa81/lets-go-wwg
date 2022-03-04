//Exercício #04
//Faça um programa que solicite à usuária que informe um número até que o número informado seja par. Seguindo o fluxo:
//Insira um número:
//> 3
//Insira um número:
//> 7
//Insira um número:
//> 2
//Agora sim podemos dividir igualmente entre mim e você!

package main

import (
	"fmt"
	"os"
)

func main() {
	var numero int

	fmt.Print("Insira um número: ")
	fmt.Fscanf(os.Stdin, "%d", &numero)

	for (numero % 2) != 0 {
		fmt.Print("Insira um número: ")
		fmt.Fscanf(os.Stdin, "%d", &numero)
	}
	fmt.Printf("Agora sim podemos dividir igulmente entre mim e você!")
}
