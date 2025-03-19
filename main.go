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

	//with an empty list as interface you can add any go datatype
	var emptyList []interface{}
	emptyList = append(emptyList, 10)
	emptyList = append(emptyList, true)
	emptyList = append(emptyList, "test")

	//iterate through the entire list and print its value
	for _, value := range emptyList {
		fmt.Println(value)
	}

	//iterate through the entire list and print its value and type if string
	for _, value := range emptyList {
		if v, ok := value.(string); ok {
			fmt.Println(v + "string")
		} else {
			fmt.Println(value)
		}
	}

	/* Generics - Slice reverse int */
	slice := []int{5, 1, 2, 3}
	newInt := reverse(slice)
	fmt.Println(newInt)

	/* Generics - Slice reverse generics */
	sliceString := []string{"a", "e", "f", "b"}
	newString := reverseGeneric(sliceString)
	fmt.Println(newString)

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

func reverse(slice []int) []int {
	newInt := make([]int, len(slice))

	newIntLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newInt[newIntLen] = slice[i]
		newIntLen--
	}
	return newInt
}

func reverseGeneric[T int | string](slice []T) []T {
	newInt := make([]T, len(slice))

	newIntLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newInt[newIntLen] = slice[i]
		newIntLen--
	}
	return newInt
}
