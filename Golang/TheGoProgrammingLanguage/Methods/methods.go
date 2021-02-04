package main

import "fmt"

type shape struct {
	height int 
	width int
}


func (s shape) Area() int{
	return s.height*s.width
}

func (s *shape) Perm() int {
	return s.height + s.width
}
func main(){

	var s shape 
	s.height = 1
	s.width = 2
	fmt.Println(s.Area())
	fmt.Println(s.Perm())
	fmt.Printf("This shapes area is %d\n", s.Area())
	fmt.Printf()
}