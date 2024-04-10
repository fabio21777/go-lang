package main

import "fmt"

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1

	return // retorno limpo
}

func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := troca(2, 3)

	fmt.Println(segundo, primeiro)
}
