//Construa uma função que receba uma lista de números inteiros,
// modifique essa lista dobrando os números ímpares e divida por dois os pares,
// e retorne a lista modificada e a soma de todos os elementos da lista.

package main

import (
	"fmt"
)

func main() {
	var numero = []int{10, 25, 30, 11, 6}
	var soma int = 0

	for x, num := range numero {
		if num%2 != 0 {
			numero[x] *= 2
		}
		if num%2 == 0 {
			numero[x] /= 2
		}

		soma += numero[x]
	}

	fmt.Printf("%v\n%d", numero, soma)
}
