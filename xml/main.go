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

	fmt.Println("\tSuccessfully open users.xml\n") //=> Print successfully

	byteValue, _ := ioutil.ReadAll(xmlFile) //=> read opened xmlFile as a byte array
	var users Users                         //=> initialize Users array
	xml.Unmarshal(byteValue, &users)        //=> byteValue -> users

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter: " + users.Users[i].Social.Twitter)
		fmt.Println("Youtube: " + users.Users[i].Social.Youtube)
		fmt.Printf("\n")
	}
}
