package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s := "😀😅testing😂"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(s[3])
	fmt.Println(s[4])
	fmt.Println(s[5])
	s += "omg "
	fmt.Println(s)

	fmt.Println(utf8.DecodeLastRuneInString("😅"))
	fmt.Println(utf8.DecodeLastRuneInString("a"))

}