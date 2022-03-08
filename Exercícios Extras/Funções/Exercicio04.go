//Crie duas funções, uma que receba uma string e retorne uma nova string alterando cada caractere para o terceiro caractere seguinte,
//e outra que reverta esse processo. Exemplo, o caractere “a” viraria “d”.

package main

import "fmt"

func main() {

	nome := "Erika"
	fmt.Println(nome)

	alterar := caractere(nome, 3)
	fmt.Println(alterar)

	fmt.Println(caractere(alterar, -3))

}
func caractere(nome string, num int) string {

	var novoNome string
	for i := 0; i < len(nome); i++ {
		novoNome += string(nome[i] + byte(num))
	}
	return (novoNome)
}
