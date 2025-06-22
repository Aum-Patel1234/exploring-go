package specifics

import (
	"fmt"
)

// CustomError is a struct that implements the error interface
type CustomError struct {
	Message string
	Code    int
}

// Implementing the error interface for CustomError
func (e CustomError) Error() string {
	// we can also do this via errors.new() from error package
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, CustomError{
			Message: "cannot divide by zero",
			Code:    400,
		}
	}
	return x / y, nil
}
