package Model

import "time"

type Compra struct {
	Itens []Item
	Date  time.Time
}

func Start(produto string) Compra {
	var listItem = []Item{}

	item := Item{Nome: produto}
	listItem = append(listItem, item)

	compra := Compra{listItem, time.Now()}

	return compra
}
