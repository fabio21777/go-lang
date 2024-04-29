package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)

	c <- 3 * base // enviando dados para o canal

	time.Sleep(time.Second)

	c <- 4 * base // enviando dados para o canal
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	fmt.Println("A") // Aparece primeiro pois a função acima é executada em uma go routine
	a, b := <-c, <-c // recebendo dados do canal, aqui o programa espera até que os dados sejam enviados e voltar a ser sincrono
	fmt.Println(a, b)

	fmt.Println(<-c)
}
