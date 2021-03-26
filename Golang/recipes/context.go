package main

import (
	"fmt"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	fmt.Println("started")
	defer fmt.Println("Server handling over")

	select {
	case <- time.After(10 * time.Second):
		fmt.Println("in here then?")
	case <- ctx.Done():
		fmt.Println("cancelled")
		err := ctx.Err()
		fmt.Println("Server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

// func main() {
// 	http.HandleFunc("/", Hello)
// 	log.Fatal(http.ListenAndServe(":8090", nil))
// }	