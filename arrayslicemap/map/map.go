package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[321] = "João"
	aprovados[456] = "Pedro"

	fmt.Println(aprovados)

	for key, nome := range aprovados {
		fmt.Printf("Nome: %s | Chave: %d \n", nome, key)
	}

	employees := map[string]float64{
		"joao Gabriel": 123.658,
		"Antonio ":     665.68,
		"Júlio":        65.218,
		"Pedro":        202.111,
	}

	fmt.Println(employees)

	employee_by_letter := map[string]map[string]float64{
		"A": {
			"Antonio Pedro": 1267.874,
			"Arlindo João":  3967.874,
		},
		"P": {
			"Paula Lima":  9967.874,
			"Perla Silva": 3568.874,
		},
	}

	fmt.Println(employee_by_letter)
}
