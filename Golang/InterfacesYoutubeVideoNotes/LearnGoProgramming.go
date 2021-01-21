package main

import (
	"fmt"
)


type book struct {
	title string
	price int
}

func (b book) print(){
	fmt.Println("Book: ", b.title)
}

type movie struct {
	title string 
	price int
}

func (m movie) print(){
	fmt.Println("Movie: ", m.title)
}

type printer interface {
	print()
}


func main(){
	var (
		mobydick = book{title:"Moby Dick", price: 10}
		harryPotter = movie{title:"Harry Potter", price: 30}
	)	
	var store list
	store = append(store, mobydick, harryPotter)
	store.print()
}