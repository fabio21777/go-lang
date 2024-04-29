package main

import "fmt"

func main() {
	ch := make(chan int, 1) // canal com buffer
	ch <- 1
	<-ch
	ch <- 2

	fmt.Println(<-ch)
}
