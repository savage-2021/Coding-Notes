package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func main() {
	// recoverFromPanic()
	// m := make(map[string]int)
	// m["key"] = 1

	// var b bytes.Buffer
	// err := json.NewEncoder(&b).Encode(&m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// raw, _ := json.Marshal(&m)
	// if string(raw) == b.String(){
	// 	fmt.Println("they're equal")
	// } else {
	// 	fmt.Printf("%s != %s", string(raw), b.String())
	// }
	// unmarshallJSON()
	// bigSlice()
	// loopTing()
	// deferExample()
	channels()
	a := escapeAnalysisStuff()
	fmt.Println(*a)
	fmt.Println(runtime.GoroutineProfile())
}


func unmarshallJSON(){
	v := []byte(`{"status": 200},{"bad": 400}`)
	var result map[string]interface{}

	if err := json.Unmarshal(v, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}


func recoverFromPanic(){
	defer func(){
		fmt.Println("revocer", recover())
		fmt.Println(recover())
	}()
	panic("not good")
}


func bigSlice(){
	b := make([]byte, 1000)
	fmt.Println(len(b), cap(b),&b[0])
	res := make([]byte, 3)
	copy(res, b[:3])
	fmt.Println(len(res), cap(res), &res[0])
}

func loopTing(){
	i := 0
	loop:
		for {
			i += 1
			if i == 20 {
				break loop
			}
		}
	fmt.Println(i)
}

func deferExample(){
	fmt.Println("Started")
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println("Done")
}


   // function 1 
    func portal1(channel1 chan string){ 
        for i := 0; i <= 3; i++{ 
            channel1 <- "Welcome to channel 1"
        } 
          
    } 
      
    // function 2 
     func portal2(channel2 chan string){ 
        channel2 <- "Welcome to channel 2"
    }

func channels(){
      
    // Creating channels 
   R1:= make(chan string) 
   R2:= make(chan string) 
     
   // calling function 1 and  
   // function 2 in goroutine 
   go portal1(R1) 
   go portal2(R2) 
     
   // the choice of selection  
   // of case is random 
   select{ 
       case op1:= <- R1: 
       fmt.Println(op1) 
       case op2:= <- R2: 
       fmt.Println(op2) 
   } 

   ch := make(chan int)
   go func(){
	ch <- 4
   }()
   select {
		case op1:= <-ch:
		fmt.Println(op1)
		default:
		fmt.Println("in default")

   }
}

func escapeAnalysisStuff() *int{
	a := 10
	return &a
}