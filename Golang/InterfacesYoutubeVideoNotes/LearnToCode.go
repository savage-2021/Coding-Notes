package main

import "fmt"

type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}

func ( s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " -- the secretAgent speak")
}

func ( p person) speak() {
	fmt.Println("I am", p.first, p.last, " -- the person speak")
}


func bar (h human) {
	fmt.Println("I was passed into bar", h)
}
var x int 
type human interface {
	speak()
}

func SwitchOnType(h interface{}) {
	switch h.(type) {
	case secretAgent:
		fmt.Println(h.(secretAgent).ltk)
	case person: 
		fmt.Println(h.(person).first)
	}
}
type hotdog int
func main() {

	sa1 := secretAgent {
		person : person{"James", "Bond"},
		ltk: true,
	}

	sa2 := secretAgent {
		person : person{"Miss", "Moneypenny"},
		ltk: true,
	}

	p1 := person {
		first: "Dr.",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}