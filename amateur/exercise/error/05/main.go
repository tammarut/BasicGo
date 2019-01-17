package main

import (
	"fmt"
	"log"
	"os"
)

//TODO: Create log_file and setoutput
//! fmt.Println =>print error terminal
//! log.Println =>print error on term log (date, time)
//! log.Fatal =>write log and "dead" exit status!
func main() {
	logs, err := os.Create("log.txt")
	defer logs.Close()
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logs)

	file, err := os.Open("anime.txt")
	defer file.Close()
	if err != nil {
		//fmt.Println("Error:", err)
		//log.Println("Error:", err)
		log.Fatalln("Fatal:", err)
	}

	fmt.Println("Everything alright.")
}

//Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.

//... the Fatal functions call os.Exit(1) after writing the log message ...
// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
