package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func commonString(s1 string) {
	hm := make(map[string]int)
	for _, el := range s1 {
		str := string(el)
		val, ok := hm[str]; if ok {
			hm[str] = val +1 
		} else {
			hm[str] = 1
		}
	}
	for key, val := range hm {
		fmt.Println(key, val)
	}
}


func createMap(s string) map[string]int32 {
	m := make(map[string]int32)
	for _, ru := range s {
		char := string(ru)
		val, ok := m[char]; if ok {
			m[char] = val + 1
		} else {
			m[char] = 1
		}
	}
	return m
}

func printMap(m map[string]int32){
	for key, val := range m {
		fmt.Printf("%s: %d\n", key, val)
	}
}

func printSortedMap(m map[string]int32){
	sorted := make([]string, len(m))
	for key := range m {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)
	for _, el := range sorted {
		fmt.Printf("%s: %d\n", el, m[el])
	}
	a := strings.Fields("tesitng don't you athinkfsn fds ")
	for _, e := range a {
		fmt.Println(e)
	}
}

// Okay so the idea is, we're going to create a channel of characters from a string 

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func(){
		defer close(intStream)
		for _, i := range integers {
			select {
			case <- done:
				return 
			case intStream <-i:
			}
		}
	}()
	return intStream
}
func generateString(done <-chan interface{}, intStream ...int) <-chan int {
	strChan := make(chan int)
	
	go func() {
		defer close(strChan)
		for i := range intStream {
			select {
			case <-done:
				return
			case strChan <- i:
			}
		}
	}()
	return strChan
}

func Go(){
	// s := "This is my test string"
	// done:= make(chan interface{})
	// defer close(done)
	// c := generateString(done, s)
	// for str := range c {
	// 	fmt.Println(str)
	// }
	done := make(chan interface{})
	defer close(done)
	intStream := generateString(done, 1,2,3,4,5)
	// pipeline := multiply(done, add(done, multiply(done, intStream,2),1),2)

	for v:= range intStream {
		fmt.Println(v)
	}
}

func main() {
	Go()

	slice := []int{1,2,3,4,5}
	arr := [5]int{}
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slice))
	for i := range slice {
		arr[i] = slice[i]
	}

}