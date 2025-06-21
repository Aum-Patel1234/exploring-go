https://go.dev/blog/declaration-syntax - why Go use syntax `var variable_name type_name`

https://go.dev/doc/effective_go#blank - for use of - "_"

# Named Return Values in Go

Named return values in Go let you declare the return variable names in the function signature. They behave like pre-declared variables at the top of the function.

A **naked return** (just `return`) will return the current values of those named variables.

✅ **Use**: For short, simple functions where return values are obvious.  
❌ **Avoid**: In long functions as it reduces readability.

---

## Example
```go
package main

import "fmt"

func calculate(x, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	return // Naked return
}

func main() {
	a, b := calculate(10, 4)
	fmt.Println("Sum:", a)
	fmt.Println("Difference:", b)
}
