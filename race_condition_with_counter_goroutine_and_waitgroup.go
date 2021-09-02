package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int
	const numGo int = 50

	for i := 1; i < numGo; i++ {
		wg.Add(1)
		go incFunc(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("Counter final:", counter)
}

func incFunc(counter *int, wg *sync.WaitGroup) {
	temp := *counter
	//Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
	runtime.Gosched()
	temp += 1
	*counter = temp
	fmt.Println("Counter: ", *counter)
	fmt.Println("Gorountines: ", runtime.NumGoroutine())
	wg.Done()
}
