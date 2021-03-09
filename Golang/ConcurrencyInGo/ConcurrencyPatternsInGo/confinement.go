package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	// adhocConfinement()
	// lexicalConfinement()
	lexicalConfinementAgain()
}

// Confienent through convention
// We can see that data is available to loopData and the loop over handleData
// -> by convention, we are only accessing it from the loopData function 
// -> people might make mistakes with this in the future, since its not explicitly said 
// -> This is why lexical confienement is betters
func adhocConfinement(){
	data := make([]int, 4) // 

	loopData := func(handleData chan<-int) {
		defer close(handleData)
		for i:= range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func lexicalConfinement(){
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // we instantiate the channel within the lexcial scope of the chanOwner function
		// This limits the scope of the write aspect of the results channel ot the closure defined below it
		// it confines the write aspect of this channel to prevent other goroutines from writing to it
		go func() {
			defer close(results)
			for i:=0; i<=5; i++{
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int){ // receive a read only viwe of the channel 
		// by declaring that the only usage we require is read only, we confine usage of the channel within the consume function
		// only to read
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving")
	}
	results := chanOwner() // we receive the read aspectt of the channel, and we're able to pass this to the consumer 
	// which can do nothing but read from it, once again confines it to a read only view of the channel
	consumer(results)
}

func lexicalConfinementAgain(){

	// Doesn't close around the data slice, needs to take it as input
	// we pass in different parts of the slice, therefore constraining the function
	// this way there is no sharing of the slice through goroutines
	// therefore no synchronising of memory
	printData := func(wg *sync.WaitGroup, data []byte){
		defer wg.Done()
		var buff bytes.Buffer 
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()

}