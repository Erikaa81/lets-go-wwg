//Crie uma função que receba uma lista de vendas de miojo, e retorne o valor vendido por vendedor com um map e o valor total vendido.
//Uma venda possui o valor, o nome da pessoa que vendeu e a data. Dica: trate o valor como inteiro.

package main

import (
	"fmt"
)

type Venda struct {
	valor        int
	nomeVendedor string
	data         string
}

func yyyymain() {

	venda1 := Venda{valor: 100, nomeVendedor: "Maria", data: "01/02"}
	venda2 := Venda{valor: 70, nomeVendedor: "João", data: "01/02"}
	venda3 := Venda{valor: 10, nomeVendedor: "Paula", data: "01/02"}

	var soma = venda1.valor + venda2.valor + venda3.valor
	{
		venda := make(map[string]int)
		venda["Maria"] = 100
		venda["João"] = 70
		venda["Paula"] = 10

		fmt.Println("map", venda)

		fmt.Println("A soma total das vendas foi:", soma)

	}
}
