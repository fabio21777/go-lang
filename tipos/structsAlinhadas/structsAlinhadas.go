package main

import "fmt"

type item struct {
	productID int
	qtd       int
	price     float64
}

type pedido struct {
	userID int
	items  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.items {
		total += item.price * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		items: []item{
			item{
				productID: 1,
				qtd:       2,
				price:     12.10,
			},
			item{
				productID: 2,
				qtd:       1,
				price:     23.49,
			},
			item{
				productID: 11,
				qtd:       100,
				price:     3.13,
			}},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
