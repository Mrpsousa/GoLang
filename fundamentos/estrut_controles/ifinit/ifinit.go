package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numero_aleatorio() int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numero_aleatorio(); i > 5 {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}
}
