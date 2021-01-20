package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	// var s, sep string
	// for i:=1; i < len(os.Args); i++{
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// This is a costly solution
	// Each time the loop goes around, the string s gets completely new cnotent s
	// The old contents of s are no loger in use, so they will be garbage-collected 
	// for _, arg := range(os.Args[1:]) {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))
}