package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	// cuidado...
	fmt.Println("Teste " + string(123)) // converte o valor 123 para o caracter correspondente na tabela Unicode

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) // converte o valor 123 para string

	// string para int
	num, _ := strconv.Atoi("123") // converte a string "123" para int
	fmt.Println(num - 122)

	// string para boolean
	b, _ := strconv.ParseBool("T") // converte a string "true" para boolean verdadeiro 'true', '1', 't' ou 'T'
	if b {
		fmt.Println("Verdadeiro")
	}

}
