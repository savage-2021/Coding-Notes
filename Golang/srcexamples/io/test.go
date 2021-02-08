package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func exampleCopy(){
	r := strings.NewReader("FUcking who knows")

	if _,err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func main(){

	exampleCopy()

}