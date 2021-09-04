package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int) // create 2 read & write channel
	c2 := make(chan int)

	go populate(c1) // populate c1 chan with data

	go fanOutIn(c1, c2) // call fanOutIn, pull data from c1 ( <-c1 ) then populate c2 (c2 <- ) afterward

	for v := range c2 { // pull data from c2 chan
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) { // populate input channel with data
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 { // pull data from c1 chan
		wg.Add(1)
		// split this work to another goroutine, 'cause timeConsumingWork() will take random amount of time to complete
		// so that if we continue this block of code, code flow is frozen for a period of time until timeConsumingWork() completed
		// by split the work to another goroutine we can countinue to create and do other jobs on other goroutines without block
		// calling goroutine
		go func(v2 int) {
			c2 <- timeConsumingWork(v2) // push data to c2 chan
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500))) // make this goroutine (or function if it is run directly in call goroutine)
	return n + rand.Intn(1000)                                   // sleep for a period of time aka wait and do nothing.
}
