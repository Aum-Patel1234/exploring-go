package main

import (
	"fmt"
	maps_example "maps_example/example"
)

func main() {
	maps_example.MapsExample()
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 3))
	maps_example.ZeroMapValueExample()
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	// m := make(map[int]int)
	for i, num := range nums {
		elem, ok := m[target-num]
		if ok {
			return []int{i, elem}
		}
		m[num] = i
	}
	return []int{}
}
