package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//! Another soulution to solve race_condition by Atomic
func main() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			r := atomic.LoadInt64(&incrementer)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counted:", incrementer)

}
