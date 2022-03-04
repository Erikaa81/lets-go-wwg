//Exercício #02
//Considerando os times do item anterior, crie uma slice para representar cada um.
//Adicione o jogador Luis Inácio no time vermelho usando a função append()
//Printe na tela os nomes dos jogadores do time vermelho
package main

import (
	"fmt"
)

func main() {
	timeAmarelo := []string{"Fernando", "Lúcia", "Mariana", "Ana"}
	fmt.Println(timeAmarelo)

	timeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}
	fmt.Println(timeVermelho)

	timeVermelho = append(timeVermelho, "Luis Inacio")
	fmt.Println("Agora os jogadores do time vermelho são: \n", timeVermelho)
}
