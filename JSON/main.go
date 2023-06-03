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
	c := cachorro{"rex", "dalmata", 3}

	dog, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dog))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poddle",
	}

	dog2, err2 := json.Marshal(c2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(bytes.NewBuffer(dog2))

}
