package interface_go

import (
	"fmt"
)

// shape is any type which implements the below 2 functions
type Shape interface {
	area() float64
	perimeter() float64
}

// rect
type Rect struct {
	width, height float64
}

func (rectange Rect) area() float64 {
	return rectange.height * rectange.width
}
func (rectange Rect) perimeter() float64 {
	return 2 * (rectange.height + rectange.width)
}

// circle
type Circle struct {
	radius float64
}

const pie = 3.14

func (c Circle) area() float64 {
	return pie * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * pie * c.radius
}

// ------------Implement-------------

func helper(shape Shape) {
	fmt.Println("The area of type", fmt.Sprintf("%T", shape), "is =", shape.area(), "and perimeter=", shape.perimeter())
}

func InterfaceExample() {
	// The area of type interface_go.circle is = 314 and perimeter= 62.800000000000004
	helper(Circle{
		radius: 10,
	})

	// The area of type interface_go.rect is = 12 and perimeter= 22.4
	helper(Rect{
		width:  10,
		height: 1.2,
	})
}
