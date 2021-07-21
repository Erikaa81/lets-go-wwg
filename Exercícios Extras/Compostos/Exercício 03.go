
package main

import ("fmt")

func main() {

	cidade := map[string]string{
		"São Paulo":     "Brasil",
		"Montevideu":    "Uruguai",
		"Roma":          "Itália",
		"Salvador":      "Brasil",
		"Florianópolis": "Brasil",
	}

	contagem := make(map[string]int)
	{

		for  indice, valor := range cidade {
			if valor == "Brasil" {
				contagem["Brasil"] = contagem["Brasil"] + 1
			} else if valor == "Itália" {
				contagem["Itália"] = contagem["Itália"] + 1
			} else if valor == "Uruguai" {
				contagem["Uruguai"] = contagem["Uruguai"] + 1
			}
			{
	
				fmt.Println(indice, valor, contagem)
			}

		}
	}
}
