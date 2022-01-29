package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorial_anterior, _ := fatorial(n - 1)
		return n * fatorial_anterior, nil
	}
}

func main() {

	valor, _ := fatorial(5)
	fmt.Println(valor)
}
