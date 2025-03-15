package main

import (
	"fmt"
	"go-to-do/GoCoursePractice/Model"
)

func main() {
	fmt.Println("Hello World")
	compra := Model.Start("item01")
	fmt.Println(compra.Date, compra.Itens[0].Nome)
}
