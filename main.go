package main

import (
	"GoPractice/Model"
	"fmt"
)

func main() {
	var listItem []Model.Item

	listItem = append(listItem, Model.Item{Name: "Item 01"})
	listItem = append(listItem, Model.Item{Name: "Item 02"})

	compra, err := Model.Start("market01", listItem)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra.Market, compra.Date, compra.Items[0].Name)
	}

	var listItemString []string

	listItemString = append(listItemString, "Item 03")
	listItemString = append(listItemString, "Item 04")

	compraString, err := Model.StartString("market02", listItemString)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compraString.Market, compraString.Date, compraString.Items[0].Name)
	}

}
