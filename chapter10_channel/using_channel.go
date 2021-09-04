package main

import (
	"fmt"
)

func main() {

	// a bidirectional channel for int with capacity of 1 int
	c := make(chan int, 1)

	go foo(c)
	bar(c)

	fmt.Printf("%p\n", &c)
	fmt.Println(c)
}

func foo(c chan<- int) {
	c <- 1
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
