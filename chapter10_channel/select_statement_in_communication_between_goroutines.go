package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sender(even, odd, quit)

	receiver(even, odd, quit) // need to set receiver on main goroutine to make
	// program wait for sender to complete

	fmt.Println("exiting")
}

func sender(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0 //send message to quit channel after iteration
}

func receiver(e, o, q <-chan int) {

	//concise select statement inside an infinite loop to make the
	//receiver alive until quit channel message received
	for {
		select {
		case v := <-e: //state that v receives/comsumes a value from e channel
			fmt.Println("from even channel:", v)
		case v := <-o:
			fmt.Println("from odd channel:", v)
		case v := <-q:
			fmt.Println("from exit channel:", v)
			return // need return to halt the receiver()
		}
	}

}
