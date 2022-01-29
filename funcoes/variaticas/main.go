package main

import "fmt"

func media(num ...float64) float64 {
	total := 0.0

	// _, ignorando indice
	for _, val := range num {
		total += val
	}

	return total / float64(len(num))

}

func main() {
	val := media(2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(val)
}
