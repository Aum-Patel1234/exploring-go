package arrays

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func VariaticFunctions() {
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 10, 20:", sum(10, 20))
	fmt.Println("Sum of 5:", sum(5))
}
