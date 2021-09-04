package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//var mutex sync.Mutex
	var wg sync.WaitGroup
	var counter int64
	const numGo int = 50

	for i := 1; i <= numGo; i++ {
		wg.Add(1)
		go incFunc(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("Counter final:", counter)
}

func incFunc(counter *int64, wg *sync.WaitGroup) {

	// mutex is used to block a certain block of code & data in memory
	// from access by any other shared memory threads/ goroutine
	//mutex.Lock()
	atomic.AddInt64(counter, 1)
	fmt.Println("Counter: ", atomic.LoadInt64(counter))
	fmt.Println("Gorountines: ", runtime.NumGoroutine())

	//mutex.Unlock()

	wg.Done()
}
