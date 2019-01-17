package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//!! Create file, Write_file and handle error
func main() {
	f, err := os.Create("movie.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	first := strings.NewReader("Thor 1")
	io.Copy(f, first)
}
