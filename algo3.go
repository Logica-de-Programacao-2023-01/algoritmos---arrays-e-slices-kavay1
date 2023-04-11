package main

import "fmt"

func main() { //Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

lista := [4]float64{1.2,3.1,4.2,5.0}
produto := 1.0

for _, valor := range lista {
	produto *= valor
}
fmt.Println(produto)



}