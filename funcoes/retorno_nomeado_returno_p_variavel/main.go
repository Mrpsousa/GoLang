package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

var soma = func(a, b int) int {
	return a + b
}

func main() {
	//retorno nomeado
	r1, r2 := trocar(1, 2)
	fmt.Println(r1, r2)

	//primeira classe
	fmt.Println(soma(1, 8))

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 3))
}
