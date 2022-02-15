package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) preco_com_descont() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var produto1 produto

	produto1 = produto{
		nome:     "Lapis",
		preco:    1.78,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.preco_com_descont())
	produto2 := produto{"notebook", 2500, 0.10}
	fmt.Println(produto2, produto2.preco_com_descont())
}
