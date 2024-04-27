package main

// os inits são executados antes da função main e na ordem em que são importados
// não pode haver dependência cíclica entre pacotes
// os nomes dos diretorios devem ser iguais aos nomes dos pacotes e dos arquivo .go
import (
	"fmt"

	"github.com/fabio21777/golang-pacotes/operations"
	// op "github.com/fabio21777/golang-pacotes/operations" // alias pode ser usado para pacotes com nomes grandes ou que sejam repetidos
	_ "github.com/fabio21777/golang-pacotes/display"
)

// execução do arquvo
// go run main.go, pos agora usa o pacote operations e main é um modulo

func main() {

	r := operations.Add(1, 2)
	s := operations.Sub(5, 3)

	fmt.Println(r)
	fmt.Println(s)
}
