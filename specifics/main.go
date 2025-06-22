package main

import (
	"fmt"
	"specifics/error_handling"
)

func main() {
	result, err := specifics.Divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	fmt.Println("Result:", result)
}
