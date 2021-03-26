package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type writer interface {
	Write([]byte) (int, error)
}

type file struct {
	data string
}

func (f *file) output(w writer) error {
	len, err := w.Write([]byte(f.data))
	if err != nil {
		return err
	} else {
		log.Println("Bytes Written: ", len)
		return nil
	}
}

type print struct{}

func (print *print) Write(p []byte) (n int, err error) {
	return fmt.Printf("%s\n", p)
}

type save struct{}

func (s *save) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile("./dat1", []byte(p), 0644)
	if err != nil {
		return -1, err
	}
	return len(p), nil
}
func main() {
	f := file{"test"}
	s := save{}
	p := print{}

	f.output(&s)
	f.output(&p)
}
