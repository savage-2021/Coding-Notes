package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boringGenerator("dan"), boringGenerator("alex"))
	for i := 0; i < 10; i++ {
		log.Println(<-c)
	}
}

func fanIn(c1, c2 <-chan string) <-chan string{
	c3 := make(chan string)
	go func(){
		c3<- <-c1
	}()
	go func(){
		c3<- <-c2
	}()
	return c3
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