package maps_example

import "fmt"

func MapsExample() {
	ages := map[string]int{
		"john":  30,
		"varun": 28,
	}

	fmt.Println(ages, "where length -", len(ages))
}

func ZeroMapValueExample() {
	m1 := make(map[int]int)
	m2 := make(map[int]string)
	m3 := make(map[int]bool)
	m4 := make(map[int]complex128)
	m5 := make(map[int]float64)

	fmt.Println("\nzero value will be given when map does not have the key - ")
	fmt.Println("for int -", m1[0])
	fmt.Println("for string -", m2[0])
	fmt.Println("for bool -", m3[0])
	fmt.Println("for complex128 -", m4[0])
	fmt.Println("for float64 -", m5[0])
}
