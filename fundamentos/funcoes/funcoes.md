# Funções

## Introdução
Em go cada pacote deve ter no menimo e no máximo um arquivo com a função main, que é o ponto de entrada do programa. As funções são blocos de código que podem ser chamados em qualquer parte do programa, e podem ou não retornar valores.Se for definido um valor de retorno é obrigatório que a função tenha um return.

```go
package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := soma(3, 10)
	imprimir(resultado)
}

```
