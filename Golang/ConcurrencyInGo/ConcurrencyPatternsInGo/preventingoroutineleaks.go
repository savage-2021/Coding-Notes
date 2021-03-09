package main

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

/*
A nil channel is passed into doWork, therefore the strings channel will
never actually get any strings written onto it, and the goroutine continuing doWork
will remain in memory for the lifetime of this process

A way to mitigate this is to establish a signal between the parent goroutine and its children
that allows the parent to signla cancellation to its children
-> By contention, this signal is usually a read-only channel named done
-> This parent goroutine passes this channel to the child goroutine and then closes this channel
	when it wants to cancel the child goroutine
Before we join two goroutines, we create a third goroutine to cancel the goroutine within doWork
after a second. This eliminates a leak
*/
func main() {
	doWork := func(strings <-chan string ) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork Exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork2 := func (
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s:= <-strings:
					fmt.Println(s)
				case <-done:
					return 
				}
			}
		}()
		return terminated
	}

	newRandStream := func() <-chan int {
		randStream:= make(chan int)
		go func(){
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				randStream <- rand.Intn(strings.NewReader(), 10)
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 Random Ints:")
	for i:= 0; i < 3; i++{
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	doWork(nil)
	done := make(chan interface{})
	stringChan := make(chan string)
	terminated := doWork2(done, stringChan)
	go func(){
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork Goroutine")
		close(done)
	}()	
	for i:= 0; i < 10; i++{
		stringChan<-"test"
	}
	<-terminated
	fmt.Println("Done.")
}