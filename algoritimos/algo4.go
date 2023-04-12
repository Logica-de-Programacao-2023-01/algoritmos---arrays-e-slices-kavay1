package main

import "fmt"

func main() { // Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores do elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
	
	var tamanho int 

	fmt.Println("Informe o tamanho da lista: ")
	fmt.Scan(&tamanho)
	lista := make([]int, tamanho)
	
	fmt.Println("Digite os números da lista: ")

	for i := 0; i < tamanho; i ++ {
		fmt.Scan(&lista[i])
	}
	
	var soma int

	for _, i := range lista {
		soma+=i
	}
	fmt.Println("Lista: ", lista)
	fmt.Println("Soma da lista: ", soma)


}