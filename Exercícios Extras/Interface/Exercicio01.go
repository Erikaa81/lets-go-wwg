//Dada a seguinte função printContent(), como poderíamos fazer com que ela funcionasse tanto para um arquivo quanto para uma string?

package main

import (
	"fmt"
	"os"
)

func printContent(a interface{}) {
	value, ok := a.(string)
	fmt.Println(value, ok)
}

func main() {

	var f1 interface {
	} = os.OpenFile
	printContent(f1)

	var f2 interface {
	} = "this is my content"
	printContent(f2)
}
