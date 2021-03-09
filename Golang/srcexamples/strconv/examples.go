package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//  stringAndInt()
	appendBool()
}

func stringAndInt() {
	i, err := strconv.Atoi("1]test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(err))
	s := strconv.Itoa(123)
	fmt.Println(reflect.TypeOf(s))

}

func appendBool(){
	b:= []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(b)
	fmt.Println(string(b))
}