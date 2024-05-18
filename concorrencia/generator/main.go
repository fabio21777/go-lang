package main

import (
	"fmt"

	"github.com/fabio21777/generator/generator"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	t1 := generator.Titulo("https://www.google.com", "https://www.youtube.com")
	t2 := generator.Titulo("https://www.coingecko.com/", "https://cryptobubbles.net/")

	t := juntar(t1, t2)

	fmt.Println(<-t)
	fmt.Println(<-t)
	fmt.Println(<-t)
	fmt.Println(<-t)
}
