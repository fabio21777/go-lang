package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println("F2:", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// é possivel retornar mais de um valor
func f5() (string, string) {
	return "Retorno1", "Retorno2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2") // é possivel chamar a função e atribuir o retorno a uma variavel
	fmt.Println(r3)
	fmt.Println(r4)
	// caso uma função retorne um valor e você não queira usar, use o underline, mas é obrigatorio ao que os dois retornos sejam usados
	// é possível tambem ignoras os retornos
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)

}
