package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

var w io.Writer



func main(){
w = os.Stdout
w = new(bytes.Buffer)
w = time.Second
}