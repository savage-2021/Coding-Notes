package main

import "fmt"

func main(){
	switch coinflip() {
	case "heads":
		fmt.Println("heads")
	case "tails":
		fmt.Println("tails")
	default: 
		fmt.Println("Landed on endge")
	}

}

func coinflip() string {
	return "heads"
}

func Signum(x int ) int {
	switch {
	case x > 0:
		return x + 1
	default:
		return 0
	case x < 0:
		return x -1
	}
}