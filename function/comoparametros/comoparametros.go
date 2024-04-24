package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func executa(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

func main() {
	resultado := executa(multiplicacao, 2, 3)
	fmt.Println(resultado)
}
