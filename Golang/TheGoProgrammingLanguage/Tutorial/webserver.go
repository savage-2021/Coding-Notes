package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int 

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "done a ting: %d", 10)
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Hits: %d", count)
	mu.Unlock()
}