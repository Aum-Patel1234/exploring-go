package main

import (
	arrays "arrays/example"
	"fmt"
)

func main() {
	fmt.Println("")
	arrays.ArraysExample()
	arrays.CustomSlice()

	arr := []int{1, 2, 3, 4, 5}

	// Deep Copy using copy()
	copyArr := make([]int, len(arr))
	copyReturnValue := copy(copyArr, arr)

	// Shallow Copy using arr[:]
	shallowArr := arr[:]

	// Modify both copied slices
	copyArr[0] = 100
	shallowArr[1] = 200

	// Print results
	fmt.Println("\nOriginal slice: ", arr)                                          // [1 200 3 4 5]
	fmt.Println("Deep copied slice: ", copyArr, "it return value", copyReturnValue) // [100 2 3 4 5]
	fmt.Println("Shallow copied slice: ", shallowArr)                               // [1 200 3 4 5]

	// Variatic funtion example
	arrays.VariaticFunctions()
}
