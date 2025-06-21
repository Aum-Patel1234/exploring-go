package structs

import "fmt"

type rect struct {
	width  float64
	height float64
}

func (obj_name rect) area() float64 { // Struct methods in Go
	return obj_name.width * obj_name.height
}

func UseStructMethod() {
	fmt.Println("\n\nStruct methods in  Go :")
	rectangle := rect{
		width:  20,
		height: 1.2,
	}
	fmt.Println("Area of rectangle with height", rectangle.height, "with width", rectangle.width, "=", rectangle.area())
}
