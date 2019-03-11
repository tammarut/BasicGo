package main

//! Practice code XML format not JSON
import (
	"fmt"
	"os"
)

func main() {
	xmlFile, err := os.Open("users.xml") //=> Open xmlFile
	if err != nil {                      //=> Handle error
		fmt.Println("Error from open xmlfile: ", err)
	}
	defer xmlFile.Close() //=> Close ด้วย

	fmt.Println("Successfully open users.xml") //=> Print successfully
}
