package main

import "fmt"

func main() {
	var notas [3]float64

	notas[0], notas[1], notas[2] = 6.5, 4.8, 9.8

	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	fmt.Println(total)

	for i := 0; i < len(notas); i++ {
		fmt.Println(notas[i])
	}

}
