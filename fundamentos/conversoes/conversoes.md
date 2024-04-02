# Conversões no go GO

## Introdução
Para fazer operações como soma multiplicação e divisão, é necessário que os tipos dos operandos sejam compatíveis. Caso contrário, é necessário fazer uma conversão de tipo.

## Conversão de tipos

a lib strconv é utilizada para fazer conversões de tipos em Go, ela possui funções para converter tipos numéricos para string e vice-versa, além de converter para boolean.

observação o casting string({meu numero}) usa a tabela unicode para converter o número em um caracter, por isso é necessário usar a função strconv.Itoa({meu numero}) para converter um número para string.


```go
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
```
