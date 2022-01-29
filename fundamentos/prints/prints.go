package main

import "fmt"

func main() {
	x := 1.2364854

	fmt.Println("o valor de x: ", x)
	fmt.Printf("o valor x tbm é: %.2f", x)

	a := 1
	b := 1.1323
	c := false
	d := "OPA"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
