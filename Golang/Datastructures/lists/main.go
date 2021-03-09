package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := make([]int, 4)
	a := copy(arr2, arr1)
	fmt.Println(arr2)
	fmt.Println(a)

	slice1 := make([]int, 10)
	for i :=0; i< 10; i++ {
		slice1[i] = i
	}
	fmt.Println(slice1)
	copy(slice1, slice1[4:])
	fmt.Println(slice1)
}