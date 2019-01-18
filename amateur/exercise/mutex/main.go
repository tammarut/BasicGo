package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	counter := 0
	const limit = 100
	wg.Add(limit)
	var mu sync.Mutex

	for i := 0; i < limit; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counted:", counter)
}

func init() {
	fmt.Println("Hi!, Solve race_condition by Mutex")
	fmt.Println("————————————————————————————————————————")
}
