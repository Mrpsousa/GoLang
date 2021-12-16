package main

import "fmt"

func imprimir_resultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}
}

func main() {
	imprimir_resultado(8)
}
