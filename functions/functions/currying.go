package funtions

import "fmt"

// The function getLogger takes a formatter function as input and returns a new function that accepts two strings, formats them using the formatter, and prints the result.
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(str1, str2 string) {
		formattedMessage := formatter(str1, str2)
		fmt.Println(formattedMessage)
	}
}

func greetingFormatter(firstName, lastName string) string {
	return "Hello, " + firstName + " " + lastName + "!"
}

func Currying() {
	// Create a logger with the greeting formatter
	greetingLogger := getLogger(greetingFormatter)

	greetingLogger("Aum", "Patel")    // Output: Hello, Aum Patel!
	greetingLogger("Harry", "Potter") // Output: Hello, Harry Potter!
}
