package main

import (
	"fmt"
	"functions/functions"
)

func calculate(x, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	return // Naked return , implicitly returns sum and diff
	// It returns the variables sum and diff declared above in func declaration
}

func getCoords() (x, y int) {
	// x and y are initialized with zero values
	return 9, 8 // this is explicit, x and y are NOT returned
}

func main() {
	a, b := calculate(10, 4)
	fmt.Println("Sum:", a)        // 14
	fmt.Println("Difference:", b) // 6

	fmt.Println(getCoords())

	result, err := funtions.Divide(10, 0)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Result - ", result)
	}

	funtions.Closure()
	funtions.Currying()
}
