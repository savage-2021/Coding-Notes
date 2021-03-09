package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("test")
	fmt.Println("test")

	pr, pw := io.Pipe()
	_, er := fmt.Fprintln(pw, "pr")
	if er != nil {
		panic(er)
	}
	io.Copy(os.Stdout, pr)
}