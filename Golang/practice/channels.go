package main


// Okay so the idea is, we're going to create a channel of characters from a string 
func generateString(done <-chan string, str string) <-chan[string] {
	strChan := make(chan string)
	defer close(strChan)
	go func() {
		for _, i := range str {
			select {
			case <-done:
				return
			default:
				strChan <- i
			}
		}
	}()
	return strChan
}

func Go(){
	s := "This is my test string"
	chan := generateString(s)
	for str := range chan {
		fmt.Println(str)
	}
}