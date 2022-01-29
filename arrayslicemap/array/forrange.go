package main

import "fmt"

func main() {
	numeros := [...]int{5, 4, 3, 2, 1}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

}
