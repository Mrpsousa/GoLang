package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}

	fmt.Println(coisa)
	coisa = 2
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "opa"
	fmt.Println(coisa2)
	coisa2 = 123
	fmt.Println(coisa2)
	coisa2 = curso{"vamos testar algo novo"}
	fmt.Println(coisa2)
}
