package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	for i < 10 {
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	// =======================================================

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("inclementando", j)
	}

	// =======================================================

	nomes := [3]string{"Biasi, Caixeta, Augusto"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// =======================================================

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// =======================================================

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Biasi",
	}
	for chave, valor := range usuario {
		fmt.Println(chave)
		fmt.Println(valor)
	}
}
