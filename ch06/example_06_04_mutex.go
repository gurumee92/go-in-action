package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("Result: ", counter)
}
