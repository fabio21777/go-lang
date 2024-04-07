package main

func notaParaConceito(n float64) {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		println("A")
	case 8, 7:
		println("B")
	case 6, 5:
		println("C")
	case 4, 3:
		println("D")
	case 2, 1, 0:
		println("E")
	default:
		println("Nota inválida")
	}
}

func main() {
	notaParaConceito(9.8)
	notaParaConceito(6.9)
	notaParaConceito(2.8)
	notaParaConceito(1.8)
	notaParaConceito(0.8)
	notaParaConceito(-10)

}
