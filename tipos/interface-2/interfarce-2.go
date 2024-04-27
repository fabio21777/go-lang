package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	f := ferrari{"F40", false, 0}
	f.ligarTurbo()

	var f2 esportivo = &f
	f2.ligarTurbo()

	fmt.Println(f, f2)
}
