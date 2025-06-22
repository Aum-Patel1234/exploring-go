package arrays

import "fmt"

func CustomSlice() {
	arr := make([]int, 5, 10)
	arr1 := make([]float64, 5)

	fmt.Println("CustomSlice example:\n", arr, "has a size of", len(arr), "and a capacity of", cap(arr))
	fmt.Println(arr1, "has a size of", len(arr1), "and a capacity of", cap(arr1))

	slice := make([]int, 2, 3)
	fmt.Println("\nBefore append:", slice, "Len:", len(slice), "Cap:", cap(slice))

	// IMPORTANT:
	// ✅ Always reassign the result of append.
	// This is the Go standard practice because the slice might move to a new memory block and also avoid diff memory writing

	slice = append(slice, 10)
	fmt.Println("After appending within capacity:", slice, "Len:", len(slice), "Cap:", cap(slice))

	slice = append(slice, 20) // This will force a new array allocation
	fmt.Println("After appending beyond capacity:", slice, "Len:", len(slice), "Cap:", cap(slice))
}
