```go
package main

import "fmt"

func main() {
	var b byte = 3 // byte é um alias para uint8

	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3 atribuição aditiva
	i -= 3 // i = i - 3 atribuição subtrativa
	i /= 2 // i = i / 2 atribuição divisiva
	i *= 2 // i = i * 2 atribuição multiplicativa
	i %= 2 // i = i % 2 atribuição modular

	fmt.Println(i)

	//multipla atribuição na mesma linha

	x, y := 1, 2

	fmt.Println(x, y)

	// troca de valores entre variáveis

	x, y = y, x
}
```
