package main

import (
	"fmt"
	"runtime"
)


func main() {
	done := make(chan interface{})
	defer close(done)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
}