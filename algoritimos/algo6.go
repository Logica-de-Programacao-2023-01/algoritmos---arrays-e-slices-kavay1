package main

import "fmt"

func main() {

	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var valor int
	fmt.Println("Digite um número: ")
	fmt.Scan(&valor)

	encontrado := false

	for _, numLista := range lista {
		if numLista == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Printf("O valor %d está presente na lista. \n", valor)
	} else {
		fmt.Printf("O valor %d não está presente na lista. \n", valor)
	}











}