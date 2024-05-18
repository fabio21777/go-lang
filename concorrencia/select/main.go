package main

import (
	"time"

	"github.com/fabio21777/generator/generator"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := generator.Titulo(url1)
	c2 := generator.Titulo(url2)
	c3 := generator.Titulo(url3)

	select {
	case <-c1:
		return url1
	case <-c2:
		return url2
	case <-c3:
		return url3

	case <-time.After(400 * time.Millisecond):
		return "Todos perderam"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.coingecko.com",
	)
	println(campeao)
}
