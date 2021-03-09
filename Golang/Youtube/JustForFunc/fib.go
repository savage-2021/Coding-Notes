package main

import "fmt"

func fib(n int) <-chan int {
	fibStream := make(chan int)
	go func(){
		defer close(fibStream)
		if n < 2 {
			fibStream <- 1
			return 
		}
		// fibStream <- fib(n-1) + fib(n-2)
		fibStream <- <-fib(n-1) + <-fib(n-2)
	}()
	return fibStream
}

func main() {
	fmt.Printf("fib(4) = %d\n", <-fib(4))
}