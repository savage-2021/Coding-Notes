package main

import "fmt"

func init(){
	fmt.Println("ran init")
}
func main() {
	a := make([]int, 5, 10)
	a[0] = 3
	printSlice("a", a)
	
	b := append(a, 1)
	b = append(b, 1)
	b = append(b, 1)
	b = append(b, 1)
	b = append(b, 1)
	b = append(b, 1) // This makes the slice's capacity double
	a[0] = 2
	printSlice("b", b)
	b = b[0: len(b) + len(a)]
	
	printSlice("b", b)
	// Hence can't be a key in a map 
	// if a == b { 
	// 	fmt.Println("yep")
	// }

	m := map[string]int {
		"UTC" : 8,
		"PCG" : 9,
	}


	if seconds, ok := m["ugh"]; ok {
		fmt.Println(seconds)
	} else {
		fmt.Println("seconds", seconds)
		fmt.Println("ok", ok)
	}

	fmt.Println(m["UTC"])
	const (
		_ = iota 
		KB = (10 * iota)
		MB 
		GB = 20 * iota
		HB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(HB)
	t := test{age: 10}
	t.increaseAge()
	fmt.Println(t.age)
	j := best(t)
	fmt.Println(j.age)
}
type test struct {
	age int
}
type best struct {
	age int
}
func (t *test) increaseAge() {
	t.age = t.age + 1
}
func (t *best) increaseAge() {
	t.age = t.age + 1
}
func (t *best) increaseAge2() {
	t.age = t.age + 2
}
func Append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data)*2))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	copy(slice[l:], data)
	return slice

}


func printSlice(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}	