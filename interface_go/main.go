package main

import (
	// "fmt"
	"fmt"
	"interface_go/interface_eg"
)

func main() {
	interface_go.InterfaceExample()

	var shape interface_go.Shape = interface_go.Circle{}
	printType(shape)
	shape = interface_go.Rect{}
	printType(shape)
}

func printType(shape interface_go.Shape) {
	if _, ok := shape.(interface_go.Circle); ok { // ok has type bool
		fmt.Println("Type is Circle")
	} else if _, ok := shape.(interface_go.Rect); ok {
		fmt.Println("Type is Rectangle")
	} else {
		fmt.Println("Type is Unknown")
	}

	switch shp := shape.(type) {
	case interface_go.Circle:
		fmt.Println("A circle", fmt.Sprintf("%T", shp))
	case interface_go.Rect:
		{
			fmt.Println("A Rectangle")
		}
	default:
		fmt.Println(shp)
	}
}
