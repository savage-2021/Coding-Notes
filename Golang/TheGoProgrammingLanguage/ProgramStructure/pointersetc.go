package main

import "fmt"

func main(){
	p := new(int) // new returns a pointer
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
		x := 1
	fmt.Printf("Address of x: %d and Value of x: %d\n", &x, x)
	fmt.Println(&x)
	fmt.Println(*&x)

}

func newInt() *int {
	return new(int)
}