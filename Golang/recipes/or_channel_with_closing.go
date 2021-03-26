package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	c1 := boring("1!")
	c2 := boring("2!")
	c3 := boring("3!")
	c4 := boring("4!")
	c5 := boring("5!")
	orChannel := or(c1, c2, c3, c4, c5)
	go func(){
		defer close(c1)
		time.Sleep(time.Second * 5)
	}()

	for {
		fmt.Println(<-orChannel)
	}

}

func or(channels ...chan string) chan string{
	switch len(channels){
	case 0: // If no channels, no elements to return
		return nil
	case 1: // if only one channel, just return that channels element
		return channels[0]
	}
	orDone := make(chan string)

	// This goroutine waits for messages on our channel without blocking
	go func(){
		defer close(orDone)
		switch len(channels){
		case 2: // This case is only put in as an optimiser
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select{
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...): // we pass the orDone channel in, so that as goroutines up the tree exit, they all do 
			}
		}
	}()

	return orDone
}

// Our channel generator function
func boring(msg string) chan string { 
    c := make(chan string)
    go func() { // We launch the goroutine from inside the function.
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
        }
    }()
    return c // Return the channel to the caller.
}