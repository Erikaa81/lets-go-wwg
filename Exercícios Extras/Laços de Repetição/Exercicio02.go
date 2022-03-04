//Exercício #02
//Considerando os tópicos que já aprendemos até agora: slices, structs ,condicionais e laços de repetição, crie um programa que traga as informações sobre apartamentos de um prédio. Passos:
//Crie uma estrutura que representa um apartamento, com campos para representar seu número, o nome da sua proprietária e se tem vaga de garagem
//Reúna as estruturas em uma slice que representa um conjunto de apartamentos
//Printe as informações de cada unidade, separando por linha, usando for range

package main

import (
	"fmt"
)

type Apartamento struct {
	numero           int
	nomeProprietária string
	vagaDeGaragem    string
}

func main() {
	apartamento1 := Apartamento{numero: 101, nomeProprietária: "Melissa", vagaDeGaragem: "tem"}
	apartamento2 := Apartamento{numero: 102, nomeProprietária: "Maria", vagaDeGaragem: "não tem"}
	apartamento3 := Apartamento{numero: 103, nomeProprietária: "Clara", vagaDeGaragem: "não tem"}
	apartamento4 := Apartamento{numero: 104, nomeProprietária: "Julia", vagaDeGaragem: "tem"}

	conjuntoDeApartamentos := []Apartamento{apartamento1, apartamento2, apartamento3, apartamento4}

	for _, apartamento := range conjuntoDeApartamentos {

		fmt.Printf("O apartamento número %d é da %s,e %s vaga de garagem! \n", apartamento.numero, apartamento.nomeProprietária, apartamento.vagaDeGaragem)
	}
}
