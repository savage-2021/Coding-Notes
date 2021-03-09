package main

// used to prepare and keep multiple instances of the same type


type Pool chan *Object

func New(total int) *Pool {
	p:= make(Pool, total)
	for i := 0; i < total; i++ {
		p <- new(Object)
	}
	return &p
}