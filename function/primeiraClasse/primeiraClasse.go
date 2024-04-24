package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	subtracao := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtracao(2, 3))
}
