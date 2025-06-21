package main

import (
	"conditionals/mathutils"
	"fmt"
	// "reflect"
)

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func main() {
	str := "aum"
	msg := fmt.Sprintf("My name is %s, and type of variable is %T and the length is %d", str, str, len(str))
	fmt.Println(msg)

	// An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.
	// if INITIAL_STATEMENT; CONDITION {
	// }
	if i := 0; i < len(str) {
		fmt.Println(str[i])
	}

	plan := "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))

	sum := conditionals.Add(5, 2)
	diff := conditionals.Subtract(10, 4)
	fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
}
