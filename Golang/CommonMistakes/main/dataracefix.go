package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		so lets have a datastructe, a map
		4 goroutines add to the map and then read from it
		then ill add a fix with channels
	*/

	m := make(map[int]string)

	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		m[1] = "one"
		wg.Done()
	}()

	// go func(){
	// 	m[2] = "two"
	// }()

	// go func(){	
	// 	m[3] = "three"
	// }()

	go func(){
		if _, ok := m[1]; !ok {
			fmt.Println("1 not there")
		}
		if _, ok := m[2]; !ok {
			fmt.Println("2 not there")
		}
		if _, ok := m[3]; !ok {
			fmt.Println("3 not there")
		}
		wg.Done()
	}()
	wg.Wait()
}