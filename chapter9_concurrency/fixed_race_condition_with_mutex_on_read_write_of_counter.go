package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var counter int
	const numGo int = 50

	for i := 1; i <= numGo; i++ {
		wg.Add(1)
		go incFunc(&counter, &mutex, &wg)
	}

	wg.Wait()
	fmt.Println("Counter final:", counter)
}

func incFunc(counter *int, mutex *sync.Mutex, wg *sync.WaitGroup) {

	// mutex is used to block a certain block of code & data in memory
	// from access by any other shared memory threads/ goroutine
	//counter here is the sharing variable between goroutines
	mutex.Lock()
	temp := *counter
	temp++
	*counter = temp
	mutex.Unlock()
	fmt.Println("Counter: ", temp)
	fmt.Println("Gorountines: ", runtime.NumGoroutine())

	wg.Done()
}
