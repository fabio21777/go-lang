package main

/*
Em go não é necessario colocar o parenteses na condição do if, porém é necessário colocar as chaves,
um outro detalhe é que é possivel adicionar os paranteses para precedencia de operações
	if nota >= 7 && (caracteristica1 && caracteristica2) {
		println("Aprovado com nota", nota)
	} else {
		println("Reprovado com nota", nota)
	}

*/
func imprimirResultado(nota float64) {
	if nota >= 7 {
		println("Aprovado com nota", nota)
	} else {
		println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
