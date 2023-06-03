// & - sempre vamos usar para pegar o endereço da memoria do ponteiro
// * - podemos usar pra duas coisas, 1: na hora de declarar o tipo e na hora de recuperar o valor de um ponteiro

package main

import "fmt"

func invert(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	invert(&number)
	fmt.Println(number)
}
