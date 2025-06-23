package main

import (
	"fmt"
	"specifics/error_handling"
)

func main() {
	str := "A😊"
	// rune is a built-in type in Go that represents a Unicode code point.
	// It is an alias for int32 and can store any Unicode character.
	// Useful when working with multi-byte characters, unlike byte (which is ASCII only).
	for i, ch := range str {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, ch, ch)
	}

	specifics.PointersExample()

	result, err := specifics.Divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	fmt.Println("Result:", result)
}
