package main

import "fmt"


func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)

	go func(){
		defer close(intStream)
		for i := range integers {
			select {
			case <-done:
				return
			default:
				intStream <-i

			}
			fmt.Printf("adding into stream %d\n", i)
		}
	}()
	return intStream
}

func generator2(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func(){
		defer close(intStream)
		for i := range integers {
			intStream <- i
		}
	}()
	return intStream
}

func main(){
	done := make(chan interface{})
	defer close(done)
	ourStream := generator(done, 1,2,3,4,5)

	for i:= range ourStream {
		fmt.Println(i)
	}

}