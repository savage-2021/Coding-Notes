package main

import (
	"bytes"
	"io"
	"os"
)

// Fprintf fdsfsdfds
func Fprintf(w io.Writer, format string, aargs ...interface{}) (int, error)

// Printf fdkslnfkdslflsd
func Printf(format string, args ...interface{}) (int, error){
	return Fprintf(os.Stdout, format, args)
}

// Sprintf nfds
func Sprintf(format string, args ...interface{})string {
	var buf bytes.Buffer 
	Fprintf(&buf, format, args...)
	return buf.String()
}

func main() {

}