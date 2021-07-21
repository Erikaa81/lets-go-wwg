//Exercício 01 Descubra por que não compila


package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}
//Não estava compilando pois estava int8
