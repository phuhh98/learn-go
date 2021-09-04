package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		value, ok := <-c
		fmt.Println("Value from channel c: ", value)
		fmt.Println("Open status: ", ok)
		if !ok {
			break
		}
	}
}
