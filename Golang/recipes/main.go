package main

import (
	"fmt"
	"recipes/concurrency"
	"time"
)

func main() {
	done := make(chan interface{})
	a := concurrency.BoringGenerator(done, "hi")
	go func(){
		time.Sleep(time.Second * 10)
		close(done)
	}()
	for b := range a {
		fmt.Println(b)
	}

	
}