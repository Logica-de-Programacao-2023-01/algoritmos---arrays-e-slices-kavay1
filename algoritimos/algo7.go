package main

import "fmt"

func main() {

	lista := []int{1, 2, 3, 4, 5}

	var num int
	fmt.Println("Digite um número inteiro: ")
	fmt.Scan(&num)

	presente := false
	for _, elementoLista := range lista {
		if num == elementoLista {
			presente = true
			break
		}
	}
	if presente == false {
		lista = append(lista, num)
	}
	fmt.Print(lista)
}
