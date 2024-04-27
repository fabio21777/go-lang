package main

import "fmt"

type curso struct {
	name string
}

func main() {
	var coisa interface{} // interface{} é um tipo que aceita qualquer tipo de dado, é um tipo genérico como o object em Java.
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = "Opa"

	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"

	fmt.Println(coisa2)

	coisa2 = true

	fmt.Println(coisa2)

}
