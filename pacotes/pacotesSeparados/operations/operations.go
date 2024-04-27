package operations

import "fmt"

func init() {
	fmt.Println("Inicializando o pacote operations")
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
