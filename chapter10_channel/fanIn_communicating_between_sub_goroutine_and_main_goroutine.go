package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go sender(even, odd)
	go receiver(even, odd, fanIn)

	for value := range fanIn { //main receiver from sub goroutine of receiver()
		fmt.Println("Main goroutine received: ", value)
	}

	fmt.Println("about to exit")

}

//Sender
func sender(even, odd chan<- int) { // sender send odd & even values over 2 channels
	for i := 0; i < 51; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

//Receiver
//receiver send back values received from sender() to fanIn channel
func receiver(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for value := range even {
			fanIn <- value
		}
		wg.Done()
	}()

	go func() {
		for value := range odd {
			fanIn <- value
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn) // close fanIn channel to halt program/ for v := range channel loop
}
