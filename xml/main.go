package main

//! Practice code XML format not JSON
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//. Array of all Users in the file
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

//. The user's struct
//. Every user have Social struct which will contain all links
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

//. Social links
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	xmlFile, err := os.Open("users.xml") //=> Open xmlFile
	if err != nil {                      //=> Handle error
		fmt.Println("Error from open xmlfile: ", err)
	}
	defer xmlFile.Close() //=> Close ด้วย

	fmt.Println("Successfully open users.xml") //=> Print successfully

	byteValue, _ := ioutil.ReadAll(xmlFile)
}
