package main

// uso de composicao

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //campo anonimo
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "f40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está como turbo ligado? %v \n", f.nome, f.turboLigado)
	fmt.Println(f)
}
