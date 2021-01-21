package main

import (
	"fmt"
	"math"
)



type shape interface {
	area() float64
	// pow() float64 // so rect and circle don't implement this, and therefore don't implement shape 
}

type circle struct {
	radius float64
}

type rect struct  {
	width float64
	height float64
}

func (r *rect) area() float64 {
	return r.height* r.width
}

func (c *circle) area() float64 {
	fmt.Println(c)
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5,7}
	shapes := []shape{&c1, &r1} // Can only use the .Area() method on this slice, since 

	for _, shape := range shapes {
		fmt.Println(shape.area()) // This is valid
		// fmt.Println(shape.width) // This is unvalid, as we can only access the interface methods
		fmt.Println(getArea(shape))
	}

}