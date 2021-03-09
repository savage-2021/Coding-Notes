package main

import "fmt"

func main() {
	a := "testing"
	b := []byte(a)
	b[0] = byte(97)
	fmt.Println(string(b))
	a = string(b)
	fmt.Println(a)
}