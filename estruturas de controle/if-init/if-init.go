package main

import (
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

/*
No Go, é possível inicializar uma variável no if, e bem similar com o for da linguagens como java e c++
a variável i só é visível dentro do bloco do if

*/

func main() {
	if i := numeroAleatorio(); i > 5 { // também suportado no switch
		println("Ganhou!!!")

	} else {
		println("Perdeu!!!")
	}

	// println(i) // não está visível
}
