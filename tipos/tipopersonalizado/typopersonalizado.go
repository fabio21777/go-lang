package main

type nota float64 // Definindo um tipo personalizado

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 6.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	case n.entre(0, 2.99):
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(6.9))
	println(notaParaConceito(2.1))
	println(notaParaConceito(11.0))
}
