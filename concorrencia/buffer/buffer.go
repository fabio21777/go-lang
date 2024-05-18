package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	c <- 1
	c <- 2
	fmt.Println("Executou")
	c <- 3
	c <- 4
	c <- 5
	c <- 6
}

func main() {
	c := make(chan int, 3) // Canal com buffer
	go rotina(c)

	time.Sleep(time.Second)
	fmt.Println(<-c)
}
