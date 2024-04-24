package main

import "fmt"

func init() { // init é uma função que é executada antes da função main cada vez que o arquivo é executado
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main")

}
