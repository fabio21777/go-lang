package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // for tradicional
		fmt.Print(i, " ")

	}

	for { // laço infinito
		println("Para sempre...")
		time.Sleep(time.Second)
	}

}
