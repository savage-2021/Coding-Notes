package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Movie struct {
		Title string `json:"title"`
		Actor string 
	}

	var movies = []Movie{
		{Title : "Catch me if you can", Actor: "That leo guy"},
		{Title: "Shrek", Actor: "No idea lol "},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", data)
}