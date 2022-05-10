package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro_1 := ferrari{modelo: "f40", turboLigado: false, velocidadeAtual: 0}
	carro_1.ligarTurbo()

	var carro_2 esportivo = &ferrari{modelo: "f40", turboLigado: false, velocidadeAtual: 0}
	carro_2.ligarTurbo()

	fmt.Println(carro_1, carro_2)
}
