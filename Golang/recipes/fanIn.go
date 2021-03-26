package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	numbers := []int{1,2,3,4}
	numberChan := make(chan int, 4)
	numberChan<- numbers[0]
	numberChan<- numbers[1]
	numberChan<- numbers[2]
	numberChan<- numbers[3]



	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	numFinders := 1 //runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
    	finders[i] = numberChan
	}
	var wg sync.WaitGroup
	wg.Add(1)
	for prime := range fanIn(done, finders...) {
    	fmt.Printf("\t%d\n", prime)
	}
	wg.Wait()

	fmt.Printf("Search took: %v", time.Since(start))
}

func fanIn(done <-chan interface{},channels ...<-chan int) <-chan int {

    multiplexedStream := make(chan int)

	// Select from all the channels
	go func(){
        for i := range channels[0] {
            multiplexedStream <- i
        }
		fmt.Println("Wait Group done")
	}()

    // Wait for all the reads to complete

    return multiplexedStream
}


func generateNumbers(numbers []int) <-chan int{
	ch := make(chan int)
	go func(){
		for _,a := range numbers{
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			ch <- a
		}
	}()
	return ch
}