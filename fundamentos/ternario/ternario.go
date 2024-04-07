package main

func obterResultado(nota float64) string {
	//return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	println(obterResultado(6.2))
	println(obterResultado(5.2))
}
