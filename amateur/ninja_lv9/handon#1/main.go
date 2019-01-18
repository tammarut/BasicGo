package main

import (
	"fmt"
	"runtime"
	"sync"
)

//! Number of Goroutine and use WaitGroup
func main() {
	var wg sync.WaitGroup

	fmt.Println("CPU cores:", runtime.NumCPU())
	fmt.Println("Start goroutines:", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("\tTokyo ghoul")
		wg.Done()
	}()
	go func() {
		fmt.Println("\tSword art online")
		wg.Done()

	}()
	fmt.Println("Middle goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("After goroutines:", runtime.NumGoroutine())
}
func init() {
	fmt.Println("Ninja lv.9: Hand-on #1")
	fmt.Println("————————————————————————————————————————")
}
