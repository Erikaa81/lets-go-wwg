//Exercício 01
//No Exercício #06 da seção "Exercícios", usamos for range para percorrer uma slice de string que representava uma lista de itens
//a comprar no mercado. Agora, resolva o mesmo exercício usando a sintaxe básica da instrução for.

package main

import "fmt"

func main() {
	listaDeMercado := []string{"abacate", "sabonete", "azeite de oliva", "tomate", "banana", "macarrão", "cebola"}

	for i := 0; i < 7; i++ {
		fmt.Println(i+1, listaDeMercado[i])
	}
}
