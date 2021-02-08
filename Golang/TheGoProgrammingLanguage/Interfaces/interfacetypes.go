package main

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type test struct {

}
func (p test) Read(a []byte) (int, error){
	return 10, nil
}

func main(){
	fmt.Println("test")
	var a test;
	var tester Reader 
	tester = a // Since type test has a Read method 
	tester.Read([]byte("fucking idk "))

	var any interface{}
	any = 1
	any = "tesitng innit"
}
