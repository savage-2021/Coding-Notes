package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	URL string = "https://swapi.dev/api/people/" 
)

func main(){
	s := time.Now()
	start()
	elapsed := time.Since(s)
	log.Printf("Requests took: %s", elapsed)
}

type Result struct {
	Name		string `json:"name"`
	Error 		bool
	ErrorMsg 	string
}

func start(){
	results := make(chan Result)
	var wg sync.WaitGroup
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			request(results, i)
		}(i)
	}
	go func(){
		wg.Wait()
		close(results)
	}()

	errorCount := 0
	errorMsgs := make([]string, 0)
	for result := range results {
		if result.Error {
			errorCount ++ 
			errorMsgs = append(errorMsgs, result.ErrorMsg)
		} else {
			log.Println(result.Name)
		}
		if errorCount > 3 {
			log.Println("\nToo many errors, breaking...")
			for _, msg := range errorMsgs {
				fmt.Printf("Error: %s\n", msg)
			}
			break
		}
	}
}


func request(results chan Result, i int){
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	api := URL + strconv.Itoa(i)
	var r Result
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		r.Error = true
		r.ErrorMsg = err.Error()
		results <- r
		return 
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		r.Error = true
		r.ErrorMsg = err.Error()
		results <-r
		return 
	}

	if res.StatusCode != http.StatusOK {
		r.Error = true
		r.ErrorMsg = strconv.Itoa(res.StatusCode)
		results <-r
		return 
	}
	defer res.Body.Close()
	
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		r.Error = true
		r.ErrorMsg = err.Error()
		results <-r
		return 
	}
	results <- r
}

