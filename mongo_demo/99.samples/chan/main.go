package main

import "log"

func main() {

	ch := make(chan int)
	//chBuffered := make(chan int)

	go func() {
		ch <- 1
	}()

	/*
	yields an additional untyped boolean result reporting whether the communication succeeded. 
	The value of ok is true if the value received was delivered by a successful send operation to the channel, 
	or false if it is a zero value generated because the channel is closed and empty.
	*/
	val, ok := <-ch
	//The multi-valued receive operation returns a received value along with an indication of whether the channel is closed.
	log.Println(val,ok)
	close(ch)
	val, ok = <-ch
	log.Println(val,ok)

}

// func isChanClosed(ch chan int) bool {

// }
