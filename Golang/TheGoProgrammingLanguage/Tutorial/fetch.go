package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	concurrentRequest()
	// for _, url := range os.Args[1] {
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	b, err := ioutil.ReadAll(resp.Body)
	// 	resp.Body.Close()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("%s", b)
	// }
}

func standardRequest(){
	for _, url := range os.Args[1:] {
		fmt.Println(url)
	}
}

func concurrentRequest(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<- ch) // receive from channel
	}

	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return 
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return 
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}


func fetch1(url string, ch chan<-string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("error")
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprint("time and bytes ", secs, nbytes)

}