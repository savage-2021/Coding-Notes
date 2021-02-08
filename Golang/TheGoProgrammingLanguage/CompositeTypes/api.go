package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type website string 

type person struct {
	Name	string `json="name"`
	Height  string `json="height` 
}

func main(){
	const api string = "https://swapi.dev/api/people/1"
	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatal("shit")
	}

	var p person
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		resp.Body.Close()
		fmt.Println(err)
	}
	fmt.Println(getDimensions())
}


func getDimensions() (height, width int, err error){
	resp, err := http.Get("http://google.com")

	if err != nil {
		return 
	}
	defer resp.Body.Close()
	height = resp.StatusCode
	width = 10
	
	return 
}


