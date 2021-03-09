package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	URL string = "https://swapi.dev/api/people/"
)

type person struct {
	Name	string `json:"name"`
}
func getRequest(i int, results chan string){
	apiCall := URL + strconv.Itoa(i)
	resp, err := http.Get(apiCall)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	var p person

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatal(err)
	} else {
		results <- p.Name
	}

}


func main(){
	results := make(chan string)
	count := 15
	var wg sync.WaitGroup
	for i := 1; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			getRequest(i, results)
		}(i)
	}
	go func(){
		wg.Wait()
		close(results)
	}()
	for i := range results {
		log.Print(i)
	}

}