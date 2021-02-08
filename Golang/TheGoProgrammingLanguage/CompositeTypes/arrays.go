package main

import (
	"fmt"
)

func main(){
	fmt.Println("started")

	months := [...]string{1: "Jan", 2: "Feb", 3:"Marc", 4: "April", 5:"May", 6:"Jun"}
	Q2 := months[1:7]
	Q3 := months[1:7]
	res := compare(Q2, Q3)
	Q3 = append(Q3, "betch")
	fmt.Println(res)
	print(Q3)
	q4 := months[2:4]
	fmt.Println(q4)
	q4 = append(q4, "ting")
	fmt.Println(q4)

	var x, y [] int
	for i:=0; i < 10; i++ {
		y = append(x, i)
		fmt.Printf("%d cap=%d\t %v\n", i, cap(y), y)
		x = y
	}
}

func compare(slice1 []string, slice2 []string) bool{
	if (len(slice1) != len(slice2)){
		return false
	}
	for i, string1 := range(slice1) {
		if(string1 != slice2[i]){
			return false
		}
	}
	return true
}

func print(slice1 []string){
	for _,val := range(slice1){
		fmt.Println(val)
	}
}

