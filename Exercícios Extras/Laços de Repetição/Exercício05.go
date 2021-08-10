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
	texto = leiaLinha()
	fmt.Println("Digite um texto:")

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
