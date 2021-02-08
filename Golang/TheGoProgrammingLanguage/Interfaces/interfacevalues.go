package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello\n"))
	fmt.Println("Done")
}