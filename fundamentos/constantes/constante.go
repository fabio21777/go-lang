package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo é inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	const e, f = 5, 6

	fmt.Println(e, f)

	g, h := 7, 8
	fmt.Println(g, h)
}
