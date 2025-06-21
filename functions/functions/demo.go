package funtions

import "errors"

func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cant divide by zero")
	}
	return x / y, nil
}
