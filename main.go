package main

import (
	"GoPractice/Model"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
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

	newString2 := reverseGenericConstraint(sliceString)
	fmt.Println(newString2)

	/* panic and defer*/

	file, err := os.Create("./settings.txt")
	//defer make the instruction at the end at the function
	defer file.Close()

	if err != nil {
		panic(err)
	}

	file.Write([]byte("teste"))

	/* go get */
	//go get github.com/fatih/color
	//only works with go.mod -> go mod init <projectname>

	printColors()

	/* go routines */

	//strconv.Itoa convert an integer for string
	//go is the same as async

	/*
		for i := 0; i < 10; i++ {
			go showMessage(strconv.Itoa(i))
		}
	*/

	//sync.WaitGroup is the same as await, but you can "group" various calls
	var wg sync.WaitGroup

	wg.Add(3)
	go callDatabase(&wg)
	go callAPI(&wg)
	go processInternal(&wg)

	wg.Wait()

	//Mutex is for lock an data for the go routine not subscribe this data
	var m sync.Mutex
	i := 0
	for x := 0; x < 10000; x++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println(i)
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("callDatabase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("callAPI")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("processInternal")
	wg.Done()
}

func showMessage(message string) {
	fmt.Println(message)
}

func printColors() {
	color.Blue("blue line")
	color.Red("red line")
	color.Magenta("And many others ..")
	color.RGB(255, 128, 0).Println("foreground orange")
	color.RGB(230, 42, 42).Println("foreground red")

	color.BgRGB(255, 128, 0).Println("background orange")
	color.BgRGB(230, 42, 42).Println("background red")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

	// Mix with RGB color codes
	color.RGB(255, 128, 0).AddBgRGB(0, 0, 0).Println("orange with black background")

	color.BgRGB(255, 128, 0).AddRGB(255, 255, 255).Println("orange background with white foreground")

	// Create SprintXxx functions to mix strings with other non-colorized strings:
	yellow := color.New(color.FgYellow).SprintFunc()
	red2 := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red2("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Printf("This %s rocks!\n", info("package"))

	// Use helper functions
	fmt.Println("This", color.RedString("warning"), "should be not neglected.")
	fmt.Printf("%v %v\n", color.GreenString("Info:"), "an important message.")

	// Windows supported too! Just don't forget to change the output to color.Output
	fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))

	// Use handy standard colors
	color.Set(color.FgYellow)

	fmt.Println("\nExisting text will now be in yellow")
	fmt.Printf("This one %s\n", "too")

	color.Unset() // Don't forget to unset

	// You can mix up parameters
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset() // Use it in your function

	fmt.Println("All text will now be bold magenta.")
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

type constraintCustom interface {
	int | string
}

func reverseGenericConstraint[T constraintCustom](slice []T) []T {
	newInt := make([]T, len(slice))

	newIntLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newInt[newIntLen] = slice[i]
		newIntLen--
	}
	return newInt
}
