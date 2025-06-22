package maps_example

import "fmt"

func MapsExample() {
	ages := map[string]int{
		"john":  30,
		"varun": 28,
	}

	fmt.Println(ages, "where length -", len(ages))
}
