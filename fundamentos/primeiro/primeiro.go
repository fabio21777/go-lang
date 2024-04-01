// programas executavaeis precisam estar no pacote main
package main

/*
Os pacotes são organizados em pastas, e para que um arquivo .go seja executado, ele precisa estar no pacote main
e é necessário declarar um ou mais imports
*/

import "fmt"

// A função main é a porta de entrada de um programa Go, não é possível ter mais de uma função main em um mesmo pacote
func main() {
	fmt.Println("Primeiro")
	fmt.Println("Programa!")
}
