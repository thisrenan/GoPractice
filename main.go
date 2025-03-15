package main

import (
	"GoPractice/Model"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	compra := Model.Start("item01")
	fmt.Println(compra.Date, compra.Itens[0].Nome)
}
