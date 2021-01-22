package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func (){
		count("sheep")
		wg.Done()
	}()

	wg.Wait()

	c := make(chan string)
	go count2("bear", c)
	
	// for {
	// 	msg, open := <- c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range c {
		fmt.Println(msg)
	}

	c1 := make(chan string, 2)
	c1 <- "hello"
	c1 <- "world"
	msg1 := <- c1
	fmt.Println(msg1)
	msg1 = <- c1
	fmt.Println(msg1)

	c2 :=make(chan string )
	c3 :=make(chan string)

	go func() {
		c2 <- "every 500ms"
		time.Sleep(time.Millisecond * 500)
	}()

	go func(){
		c3 <- "every 2 seconds"
		time.Sleep(time.Second * 2)
	}()
	
	for {
		select {
		case msg2 := <- c2:
			fmt.Println(msg2)

		case msg3 := <- c3:
			fmt.Println(msg3)
		}
		// fmt.Println(<- c2) // channel 2 and channel 3 are blocking each other
		// fmt.Println(<- c3)
	}
}

func worker(jobs <-chan int, results chan<-int){
	for n := range jobs {
		results <- fib(n)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n 
	}
	return fib(n-1) + fib(n-2)
}
func count(thing string){
	for i := 1; i< 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond*500)
	}
}

func count2(thing string, c chan string){
	for i := 0; i < 5; i++{
		c <- thing
		time.Sleep(time.Millisecond *500)
	}
	close(c)
}