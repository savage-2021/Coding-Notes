package main

import "fmt"

func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func(){
		defer close(valStream)
		for {
			select {
			case <-done:
				return 
			case v, ok := <-c:
				if ok == false {
					return 
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
func main(){
	done := make(chan interface{})
	defer close(done)
	myChan := make(chan interface{})
	for val := range orDone(done, myChan){
		fmt.Println(val)
	}
}