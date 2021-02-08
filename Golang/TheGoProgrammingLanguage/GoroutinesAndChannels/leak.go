package main

import (
	"fmt"
	"net/http"
)

func main() {
	responses := make(chan string)
	go func() { responses <- http.Get("http://www.asia.gopl.io").Body() }()
	go func() { responses <- "europe.gopl.io" }()
	go func() { responses <- "americas.gopl.io" }()
	fmt.Println(<-responses) // return the quickest response
}