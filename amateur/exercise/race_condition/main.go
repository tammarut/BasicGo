package main

import (
	"fmt"
	"runtime"
	"sync"
)

//! Race condition, try to run this code below.
func main() {
	fmt.Println("CPU core:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup //TODO: 0
	counter := 0
	const gs = 100
	wg.Add(gs) //TODO: 1

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done() //TODO: 2
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait() //TODO: 3
	fmt.Println("Counted:", counter)

}
func init() {
	fmt.Println("Hi!, Oh! race condition.")
	fmt.Println("————————————————————————————————————————")
}
