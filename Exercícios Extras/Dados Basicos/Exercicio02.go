
//Exercício #02 
//Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
// Como podemos pegar a informação do ano diretamente do console?


package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	nascimento := 1950
	idade := t.Year() - nascimento

	fmt.Scan("%s", &t)
	fmt.Scan("%v", &nascimento)
	fmt.Scan("%d", &idade)

	fmt.Printf("Hoje %s a pessoa que nasceu em %v tem %d anos.",
		t, nascimento, idade)
}
