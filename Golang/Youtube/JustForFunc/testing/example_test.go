package sum

import "fmt"

// Ints returns the sum of the list of numbers
func Ints(vs... int ) int {
	return ints(vs...)
}

func ints(vs ...int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]...) + vs[0] 
}

func ExampleInts(){
	s := Ints(1,2,3,4,5)
	fmt.Println("Sum of one to 5 is 15", s)
	// Output:
	// 15

}