package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 4)
	go func() {
		c <- 1234
	}()

	fmt.Println(<-c)
}
