package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// var a rune = '日'
	// var b byte = 'X'
	// fmt.Printf("%x \n", a)
	// fmt.Println(b)
	buffer()
}

func compare(){
	var a []byte = []byte{10, 32, 42}
	var b []byte = []byte{97, 98, 99, 100}
	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(bytes.Compare(a, b))
}

func buffer(){
	// a buffer is a variable sized buffer of bytes with read and write methods
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Println(&b)
	b.WriteTo(os.Stdout)
	// A buffer can turn a string or a []byte into a reader 
	// b.Reset()
	be := bytes.NewBuffer([]byte("Hello"))
	io.Copy(os.Stdout, be)
	fmt.Println()

	buf := bytes.NewBufferString(string([]byte{97, 98, 99}))
	fmt.Println(buf)
	buf.Write([]byte{100})
	fmt.Println(buf)
	buf.Write([]byte{101})
	fmt.Println(buf)
	fmt.Println(buf.Read([]byte{100,0, 0}))
	fmt.Println(buf)
	fmt.Println(&buf)

	// Cool, so a bytes.Buffer bascially can take in a byte array, and can read and write from it nicely 
}