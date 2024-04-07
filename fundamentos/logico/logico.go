package main

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	println("dois trabalhos tv 50", tv50, tv32, sorvete)
	tv50, tv32, sorvete = compras(true, false)
	println("um trabalho tv 32", tv50, tv32, sorvete)
	tv50, tv32, sorvete = compras(false, false)
	println("nenhum trabalho ", tv50, tv32, sorvete)
}
