package main

import "fmt"

func imprimi_aprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main() {
	aprovados := []string{"Maria", "Pedro", "João", "Rafael", "Jorge"}

	imprimi_aprovados(aprovados...)
}
