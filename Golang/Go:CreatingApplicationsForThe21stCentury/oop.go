package main

import "fmt"

type specialitem struct {
	Item item 
	markup float64
}

type item struct {
	id       string
	price    float64
	quantity int
}

func (item *item) Cost() float64 {
	return item.price * float64(item.quantity)
}

func(item *specialitem) Cost() float64 {
	return item.Item.Cost() * item.markup
}


func main() {
	item := item{"1", 10, 1}
	specialitem := specialitem{item, 10}
	fmt.Println(specialitem.Cost())
}