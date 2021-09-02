package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {

	exercise for concurency using WaitGroup
	fmt.Println("started")
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	go foo(1, &wg)
	go foo(2, &wg)

	wg.Wait()
	fmt.Println("exit")
}

func foo(a int, wg *sync.WaitGroup) {
	fmt.Println("foo", a)
	fmt.Println(runtime.NumGoroutine())
	wg.Done()
}
