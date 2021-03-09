package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const url = "https://swapi.dev/api/people/1"

func main() {
	
	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Close = true
	var transport = &http.Transport{DisableKeepAlives: true}
	var myClient = &http.Client{Timeout: 10 * time.Second, Transport: transport}
	resp, err := myClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}
	type person struct {
		Name string `json:"name"`
	}
	p := &person{}
	err = json.NewDecoder(resp.Body).Decode(p)
	fmt.Println(p.Name)
	// resp, err := http.Get("")
	// // if you don't read the response body, you still need to close it
	// if resp != nil {
	// 	defer resp.Body.Close()
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = io.Copy(ioutil.Discard, resp.Body)  
	// type person struct {
	// 	Name string `json:"name"`
	// }
	// p := &person{}
	// err = json.NewDecoder(resp.Body).Decode(p)
	
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(body))
	
	// fmt.Println(p)
}