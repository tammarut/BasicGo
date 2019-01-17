package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Also handle error
//! Open file and read inside.
func main() {
	path := "/User/Admins/Go/src/github.com/tammarut/Together/Learning/Pleum/exercise/error/movie.txt"

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
