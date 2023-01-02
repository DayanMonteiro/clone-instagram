package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)
	// o pacote bytes.NewBuffer converter o json para um formato sem sem em bytes como no print acima
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	fmt.Println("Mensagem 2")

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

	c3 := map[string]string{
		"nome": "Layla",
		"raca": "Lhasa apso",
	}

	fmt.Println("Mensagem do cachorro 3")

	cachorro3EmJSON, erro := json.Marshal(c3)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro3EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro3EmJSON))

}
