package main

func inc1(n int) int {
	return n + 1
}

// função usando ponteiro

func inc2(n *int) {
	*n++
}

func main() {
	n := 1
	inc1(n)    // por valor
	println(n) // 1 não muda o valor de n
	inc2(&n)   // por referência
	println(n) // 2 muda o valor de n
}
