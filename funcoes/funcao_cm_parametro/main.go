package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

//exec recebe uma 'funcao' que é do tipo 'func(int, int) int' passamos + 2 paramets p1 e p2 que são do tipo int e exec return um int
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
func main() {
	resultado := exec(mult, 2, 4)

	fmt.Println(resultado)
}
