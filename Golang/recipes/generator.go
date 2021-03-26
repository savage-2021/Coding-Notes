package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
	1) Common in Go to have a for select loop which loops infinitely until stopped
	2) Goroutines = unit of work which can run parallel to other units
	3) A Generator is any function that takes a list of discrete values are turns them into streams
*/

// Creating a generator which gives random numbers

func main(){
	done := make(chan interface{})
	rands := randStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-rands)
	}
	close(done)

	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}

func randStream(done <-chan interface{}) <-chan int{
	rands := make(chan int)
	go func(){
		defer log.Println("Rand Stream Exited and Cleaning Up Goroutine")
		defer close(rands)
		for {
			select {
			case rands <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return rands
}


