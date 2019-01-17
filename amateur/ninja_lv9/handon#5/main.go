package main

import (
	"fmt"
	"runtime"
)

//! Try command [gobuild, ./tap]  [goinstall]
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
