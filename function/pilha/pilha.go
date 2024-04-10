package main

import "runtime/debug"

/*

Pilha de execução

A pilha funcionar de forma similar nas outras linguagens, cada função que é chamada é empilhada e quando a função termina ela é desempilhada.
*/

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
