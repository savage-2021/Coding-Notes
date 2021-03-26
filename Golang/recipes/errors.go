package main

import "fmt"

type myError struct {
	msg string
}

func(m *myError) Error() string {
	return m.msg
}
func main(){
	if err := do(); err != nil {
		fmt.Println(err)
	}
}

func do() error {
	return &myError{"Fucked the error"}
}