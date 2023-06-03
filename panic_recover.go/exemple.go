package panicrecovergo

import "fmt"

func divide(a, b float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ocorreu um erro:", r)
		}
	}()
	if b == 0 {
		panic("divisão por zero")
	}
	return a / b
}

func main() {
	resultado := divide(10, 2)
	fmt.Println("Resultado:", resultado)

	resultado = divide(10, 0)
	fmt.Println("Resultado:", resultado)
}
