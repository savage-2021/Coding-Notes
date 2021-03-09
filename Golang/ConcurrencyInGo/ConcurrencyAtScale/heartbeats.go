package main

import "time"

func doWork(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time){
	heartbeat := make(chan interface{})
	results := make(chan time.Time)
	go func(){
		defer close(heartbeat)
		defer close(results)

		pulse := time.Tick(pulseInterval)
		workGen := time.Tick((2*pulseInterval))

		sendPulse := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				}
			}
		}
	}()
}

func main(){

}