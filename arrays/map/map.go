package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678979] = "Pedro"
	aprovados[12345678980] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println("--->", aprovados[12345678978])

	// inicializando o map
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}
	// adicionando um novo elemento
	funcsESalarios["Rafael Filho"] = 1350.0

	// deletando um elemento inexistente isso não gera erro
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	// mapas aninhados

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga Pereira":  8456.78,
		},
		"F": {
			"Filipe Ribeiro":   11325.45,
			"Fabio Souza":      15456.78,
			"Fernanda Miranda": 1200.00,
		},
		"J": {
			"José João":  11325.45,
			"João Pedro": 1200.00,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	delete(funcsPorLetra, "J")

	fmt.Println(funcsPorLetra)

	// acessando map dentro de map
	fmt.Println("Acessando map dentro de map")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
