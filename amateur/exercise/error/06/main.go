package main

import (
	"fmt"
	"os"
)

//! log.Fatal => exited don't run anything continuely.
//! log.Panic => print out log to terminal and panic.
//! panic => panic but don't die yet.
func main() {
	//defer foo()
	_, err := os.Open("anime.txt")
	if err != nil {
		//log.Fatalln("Error:", err)
		//log.Panicln("Error:", err)
		panic(err)
	}
}
func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

//.. the Fatal functions call os.Exit(1) after writing the log message ...
// Fatalln = Println() followed by a call to os.Exit(1).

//Panicln = Println() followed by a call to panic().
