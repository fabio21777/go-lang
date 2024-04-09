package main

import "fmt"

func main() {
	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array.
	s2 := a2[1:3]

	fmt.Println(a2, s2)

	s3 := a2[:2] // Novo slice, mas aponta para o mesmo array

	fmt.Println(a2, s3)

	// Você pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]

	fmt.Println(s2, s4)

	//criando um slice com make (mais comum)
	s5 := make([]int, 10)

	s5[9] = 12

	fmt.Println(s5)
	// inicializa com a capacidade de 20,tamanho do array de referencia
	s5 = make([]int, 10, 20)
	// cap é a capacidade do slice
	fmt.Println(s5, len(s5), cap(s5))
	// é possivel adicionar elementos ao slice
	s5 = append(s5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s5, len(s5), cap(s5))
	// se adicionar mais elementos que a capacidade do slice, ele dobra a capacidade, o comportamento é dobrar a capacidade
	s5 = append(s5, 1)

	fmt.Println(s5, len(s5), cap(s5))
}
