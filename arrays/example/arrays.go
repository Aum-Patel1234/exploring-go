package arrays

import "fmt"

func ArraysExample() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)
	slice := arr[1:4] // here 4 is exclusive

	fmt.Println("The type of slice is -", fmt.Sprintf("%T", slice), "whereas the type of arr:", fmt.Sprintf("%T", arr), "and it contains - ", slice)
	// The type of slice is - []int whereas the type of arr: [8]int and it contains -  [2 3 4]

	// NOTE: important conversion
	fmt.Println("The type conversion from array of fixed size to slice is (arr[:]) - ", fmt.Sprintf("%T", arr[:]))
}
