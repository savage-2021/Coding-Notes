package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	// Boring
	// c := make(chan string)
	// go boring("Boring", c)
	// for i := 0; i < 5; i ++ {
	// 	fmt.Printf("You say: %q\n", <-c)
	// }
	// Boring Generator
	ch1 := boringGenerator("Alex")
	ch2 := boringGenerator("Dan")
	ch3 := fanIn(ch1, ch2)
	for i := 0; i < 15; i++ {
		fmt.Println(<-ch3)
	}
}

func boring(msg string, c chan string){
	for i := 0; ; i++{
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
	}
}

func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	go func(){
		for i := 0; ;i++{
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(ch1, ch2 <-chan string) <-chan string{
	ch3 := make(chan string)
	go func(){
		for {
			ch3<- <-ch1
		}
	}()
	
	go func(){
		for {
			ch3<- <-ch2
		}
	}()
	return ch3
}