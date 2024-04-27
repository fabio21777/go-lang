package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct {
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func (b bmw7) ligarTurbo() {
	fmt.Println("Ligando turbo...")
}

func main() {
	var b esportivoLuxuoso = bmw7{} //
	b.fazerBaliza()
	b.ligarTurbo()

}
