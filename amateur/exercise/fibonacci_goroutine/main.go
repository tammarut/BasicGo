package main

import "fmt"

func fibonacci(chNo chan int, chQ chan struct{}) {

	a, b := 0, 1

	for {
		select {
		case chNo <- a:
			a, b = b, a+b
		case <-chQ:
			return
		}
	}
}

func main() {
	//* create 2 channels
	chNo := make(chan int)
	//? what's this below? struct for what!?
	chQ := make(chan struct{})
	go fibonacci(chNo, chQ)
	for i := 0; i <= 10; i++ {
		fmt.Println(<-chNo)
	}
	//? this's below one
	chQ <- struct{}{}
}

//! Run init first before main
func init() {
	fmt.Println("Hi, fibonacci goroutine and channel.")
	fmt.Println("-----------------------------")
}
