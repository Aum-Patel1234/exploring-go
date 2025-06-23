package specifics

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c *Circle) assignRadius(newRadius float64) {
	c.radius = newRadius
}

func PointersExample() {
	circle := Circle{radius: 7}
	fmt.Println("Circle has radius -", circle.radius)
	circle.assignRadius(10)
	fmt.Println("Circle has radius -", circle.radius, "after changing via pointer.")
}
