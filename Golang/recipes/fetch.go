package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "https://swapi.dev/api/people/1"

func main() {
	fetch()
	fetchWithContext()
}


func fetch(){
	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Close = true // closes the request once the request has been made

	// Transport caches connections for future use 
	var transport = &http.Transport{DisableKeepAlives: true}
	var myClient = &http.Client{Timeout: 10 * time.Second, Transport: transport}
	resp, err := myClient.Do(req)
	
	// Best way to close the resp body
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Println(err)
	}

	type person struct {
		Name string `json:"name"`
	}

	// Encoder 
	p1 := &person{}
	err = json.NewDecoder(resp.Body).Decode(p1)

	if err != nil {
		log.Println(err)
	}

	// Marshall
	/*
		req.Body is already an io.Reader. 
		Reading its entire contents and then performing json.Unmarshal wastes resources if the stream was, say a 10MB block of invalid JSON. 
		Parsing the request body, with json.Decoder, as it streams in would trigger an early parse error if invalid JSON was encountered. 
		Processing I/O streams in realtime is the preferred go-way.
	*/
	body, err := ioutil.ReadAll(resp.Body) // reqest body is empty from above, so this p2 is empty
	p2 := &person{}
	err = json.Unmarshal(body, &p2)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(p1.Name)
	fmt.Println(p2.Name)
	return
}

func fetchWithContext(){
	
	ctx, cancel := context.WithCancel(context.Background())

	req, err := http.NewRequestWithContext(ctx, "GET", "http://httpbin.org/delay/3", nil) 
	if err != nil {
		log.Println(err)
	}
	req.Close = true
	var transport = &http.Transport{DisableKeepAlives: true}
	var myClient = &http.Client{Timeout: 10 * time.Second, Transport: transport}
	
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	resp, err := myClient.Do(req)
	p := []byte{}
	a, err := resp.Body.Read(p)
	fmt.Println(a)

	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println(err)
	}

}