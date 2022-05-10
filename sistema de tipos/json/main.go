package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	//struct to json
	p1 := produto{1, "Notebook", 1990.98, []string{"Promoção, Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json to struct
	var p2 produto
	jsonString := `{"id":2,"name":"Caneta","preco":200.13,"tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)

}
