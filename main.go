package main

import (
	"GoPractice/Model"
	"errors"
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

	/* template test */

	rectangle := Model.Rectangle{
		Width:  1,
		Height: 2,
	}

	Model.ShowRectangleArea(rectangle)

	circle := Model.Circle{
		Radius: 1,
	}

	Model.ShowCircleArea(circle)

	Model.ShowGeometry(rectangle)
	Model.ShowGeometry(circle)

	/* Error tests */

	ShowError(errors.New("a Error"))

	p := NetWorkProblem{
		NetWork:  true,
		Hardware: false,
	}

	ShowError(p)

	/* Empty interfaces tests */
	var emptyInterface interface{}
	emptyInterface = Model.Circle{Radius: 10}
	fmt.Println(emptyInterface)

}

type NetWorkProblem struct {
	NetWork  bool
	Hardware bool
}

func (p NetWorkProblem) Error() string {
	if p.NetWork {
		return "Network problem"
	} else if p.Hardware {
		return "Hardware problem"
	} else {
		return "Another Problem"
	}

}

func ShowError(err error) {
	fmt.Println(err.Error())
}
