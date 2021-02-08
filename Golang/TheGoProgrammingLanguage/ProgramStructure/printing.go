package main

import "fmt"

const boilinfF = 212.09898358492357489237504837502934075928437


func main(){
	var f = boilinfF
	var c = (f-32)*5 / 9
	fmt.Printf("boiling point = %g or %g\n", f, c)
	fmt.Printf("boiling point = %.2f or %.3f\n", f, c)

	a,b,c := 1,2,3
	fmt.Println(a,b,c)


}

func returnAddress(x int) *int {
	return &x
}

func incr(p *int) int {
	*p++
	return *p
}