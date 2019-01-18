package main

import (
	"fmt"
	"runtime"
	"sync"
)

//! Goroutines incrementer make race_condition | solve by Mutex
func main() {
	fmt.Println("Start goroutine:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mu sync.Mutex
	incrementer := 0
	const gs = 100

	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Middle goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counted:", incrementer)
}
