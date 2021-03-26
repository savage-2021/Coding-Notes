package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// BoringGenerator Genertor is just a function that returns a channel
// This jsut returns the msg forever at random times
func BoringGenerator(done chan interface{}, msg string) <-chan string { 
	c := make(chan string)
	go func(){
		for {
			select{
			case <-done:
				close(c)
				return
			default:
				c <- fmt.Sprintf("%s", msg)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}
	}()
	return c
}