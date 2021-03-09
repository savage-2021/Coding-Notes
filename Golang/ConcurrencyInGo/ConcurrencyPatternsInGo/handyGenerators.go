package main

import (
	"fmt"
	"math/rand"
)

// Repeats the values you pass it indefinitely until you tell it to stop
func repeatGenerator(done <-chan interface{}, values ...interface{}) <-chan interface{} {
    valueStream := make(chan interface{})
    go func() {
        defer close(valueStream)
        for {
            for _, v := range values {
                select {
                case <-done:
                    return
                case valueStream <- v:
                }
            }
        }
    }()
    return valueStream
}

// Takes only the first num items off its incoming valueStream and then exits 
func takeGenerator(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
    takeStream := make(chan interface{})
    go func() {
        defer close(takeStream)
        for i := 0; i < num; i++ {
            select {
            case <-done:
                return
            case takeStream <- <- valueStream:
            }
        }
    }()
	return takeStream
}

func funcGenerator(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func (){
		defer close(valueStream)
		for {
			select {
			case <-done:
				return 
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func(){
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return 
			case stringStream <- v.(string):
				fmt.Println("in case")
			default:
				fmt.Println("in default")

				stringStream <- v.(string)
			}
		}
	}()
	return stringStream
}

func exampleUseGeneratorFunction(){
	done := make(chan interface{})
	defer close(done)
	rand := func() interface{} {
		return rand.Int()
	}
	for num:= range takeGenerator(done, funcGenerator(done, rand), 10) {
		fmt.Println(num)
	}
}
func main(){
	done := make(chan interface{})
	defer close(done)
	/*
		Repeat and take generator are very efficient, although we have the capability of generating an infinite stream of ones,
		we only generate N+1 intstances, where N is the number we pass into the take stage
	*/
	// for num := range takeGenerator(done, repeatGenerator(done,1), 10) {
	// 	fmt.Printf("%v ", num)
	// }

	var message string
	for token := range toString(done, takeGenerator(done, repeatGenerator(done, "I", "am."), 5)) {
    	message += token
	}

	fmt.Printf("message: %s...", message)
}

