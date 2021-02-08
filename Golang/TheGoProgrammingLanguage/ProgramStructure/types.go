package main

import "fmt"

// Ngl, this can make things like os mych more readble
// Named types allow you to define more behaviours on a given type

type Celsius float64 
type Fahrenheit float64 

func main(){
	const (
		AbsoluteZeroC Celsius = -273.15 
		FreezingC     Celsius = 0
		BoilingC      Celsius = 100
	)
	fmt.Println(AbsoluteZeroC.String())
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c)
}

// For example, you can do that here
func(c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}