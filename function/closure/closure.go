package main

import "fmt"

func closure() func() {
	x := 20
	// Closure
	imprimeX := func() {
		fmt.Println(x)
	}
	return imprimeX
}

func main() {
	x := 10
	fmt.Println(x)

	imprimeX := closure()

	imprimeX()
}
