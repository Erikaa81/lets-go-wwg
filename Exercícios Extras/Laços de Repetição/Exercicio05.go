package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func leiaLinha() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
func main() {
	var texto string
	fmt.Println("Digite um texto:")
	texto = leiaLinha()

	letras := make(map[string]int)

	for _, letra := range texto {
		letras[string(letra)] += 1
	}

	ordenar := make([]string, len(letras))
	for i := range letras {
		ordenar = append(ordenar, i)
	}
	sort.Strings(ordenar)

	for _, l := range ordenar {
		fmt.Printf("%s => %v ", l, letras[l])
	}
}
