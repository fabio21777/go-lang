package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (tamanho fixo)
	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	fmt.Println(notas)

	//notas[3] = 10 // índice inválido

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f\n", media)

	// for range

	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	//sem usar o índice
	for _, num := range numeros {
		fmt.Printf("%d ", num)
	}
}
