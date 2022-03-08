//Dada a seguinte função, o que devemos escrever para que ela printe “inteiro” caso o parâmetro recebido seja int32 e
//“ponto flutuante” caso o parâmetro seja float32?
//Caso o tipo não seja reconhecido, a função deverá retornar um erro informando o ocorrido.

package main

import "fmt"

func printType(i interface{}) {
	switch i.(type) {

	case int32:
		fmt.Println("Tipo: int, Valor:", i.(int32))
	case float32:
		fmt.Println("\nTipo: float64, Valor: ", i.(float32))
	default:
		fmt.Println("\nTipo não encontrado")
	}

}

func main() {
	printType(int32(2))
	printType(float32(2.23))
	printType(int(2))
}
