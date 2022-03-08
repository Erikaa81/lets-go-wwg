//Construa uma função que receba uma palavra ou frase e uma letra, e retorne o número de ocorrências da letra informada.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var frase = "Estudar é preciso"
	var letra = "r"
	var ocorrencia = contarLetra(frase, letra)

	fmt.Printf(" A letra %s foi encontrada %d vezes na frase '%s'", letra, ocorrencia, frase)
}

func contarLetra(frase string, letra string) int {

	var final = strings.Count(frase, letra)

	return final
}
