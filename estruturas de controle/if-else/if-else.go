package main

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 7 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(7.2))
	println(notaParaConceito(5.5))
	println(notaParaConceito(3.3))
	println(notaParaConceito(1.2))
}
