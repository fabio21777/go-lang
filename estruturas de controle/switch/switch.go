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

func notaParaConceitoswitch2(n float64) {
	switch {
	case n >= 9 && n <= 10:
		println("A")
	case n >= 7 && n < 9:
		println("B")
	case n >= 5 && n < 7:
		println("C")
	case n >= 3 && n < 5:
		println("D")
	case n >= 0 && n < 3:
		println("E")
	default:
		println("Nota inválida")
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	//switch com valor
	notaParaConceito(9.8)
	notaParaConceito(6.9)
	notaParaConceito(2.8)
	notaParaConceito(1.8)
	notaParaConceito(0.8)
	notaParaConceito(-10)

	//switch executa primeiro case que for verdadeiro
	notaParaConceitoswitch2(9.8)
	notaParaConceitoswitch2(6.9)
	notaParaConceitoswitch2(2.8)
	notaParaConceitoswitch2(1.8)
	notaParaConceitoswitch2(0.8)
	notaParaConceitoswitch2(-10)

	//switch com tipo

	println(tipo(2.3))
	println(tipo(1))
	println(tipo("Opa"))
	println(tipo(func() {}))
	println(tipo(true))
}
