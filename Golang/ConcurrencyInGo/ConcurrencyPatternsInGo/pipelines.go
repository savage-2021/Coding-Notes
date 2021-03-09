package main

import "fmt"

/*
	Reduce memory functions, however, limits the heavy lifting of our functions
	relies on range doing heavy lifting of our pipeline, and also

*/
func main() {
	multiply := func(values []int, multiplier int) []int {
		// multipliedValues := make([]int, len(values))
		for i, v := range values {
			values[i] = v * multiplier
		}
		return values
	}

	add := func(values []int, adder int) []int {
		addedValues := make([]int, len(values))
		for i, val := range values {
			addedValues[i] = val + adder
		}
		return addedValues
	}
	vals := []int{1,2,3,4,5}
	a := add(multiply(vals, 5), 5)
	for _, i := range a {
		fmt.Println(i)
	}
	for _, i := range add(multiply(vals, 4), 3) {
		fmt.Println(i)
	}
}