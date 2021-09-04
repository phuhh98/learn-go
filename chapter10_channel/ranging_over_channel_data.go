package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // problem solved when c channel is closed
	}()

	//receiver
	for v := range c { // v is iterate over the range of receive channel c
		// v is always waiting for the next message
		// so that main func is freeze at the statement to iterate over range of c
		// unless c is closed (gone), main func will freeze and cause deadlock -- wait and do nothing
		fmt.Println(v)
	}

	fmt.Println("exiting")
}
