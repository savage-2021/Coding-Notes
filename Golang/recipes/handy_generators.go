package main

import (
	"fmt"
	"math/rand"
)

func main(){
	efficientWayOfGeneratingAmount()
	generateFunctionWhichCreatesRandomNumbers()
}
// Repeat generator blocks on the takes receive, its a super efficient
// We have the capability of generating infinite 1s, but we only generate n+1, where n is how many we want
func efficientWayOfGeneratingAmount(){
	done := make(chan interface{})
	defer close(done)
	for num := range take(done, repeat(done, 1), 10) {
    	fmt.Printf("%v", num)
	}
	fmt.Println()
}

func generateFunctionWhichCreatesRandomNumbers(){
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Int()}

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

// Repeat generator takes any value you pass into it and returns it forever
func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{}{
	valueStream := make(chan interface{})
	go func(){
		defer close(valueStream)
		for {
			for _,a := range values{
				select {
				case <-done:
					return
				case valueStream<-a:
				}
			}
		}
	}()
	return valueStream
}

// Repeat Generator that repeasts calls any function
func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func(){
		defer close(valueStream)
		// defer log.Println("broek loop")
		for {
			select {
			case <-done:
				return
			case valueStream<-fn():
			}
		}

	}()
	return valueStream
}


// Take generator only takes the first n items off a stream 
func take(done <-chan interface{}, valueStream <-chan interface{}, n int) <-chan interface{}{
	takeStream := make(chan interface{})
	go func(){
		defer close(takeStream)
		for i :=0; i < n; i++ {
			select {
			case <-done:
				return
			case takeStream<- <-valueStream:
			}
		}
	}()
	return takeStream
}
