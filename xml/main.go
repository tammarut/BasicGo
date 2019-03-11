package main

//! Example2 parse XML format
//=> XML => JSON
//? JSON => XML
import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

//. Array of all Data
type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"`
	PersonList []Person `xml:"person" json:"person"`
}

//. The person's struct
type Person struct {
	FirstName string   `xml:"firstname" json:"firstname,omitempty"`
	LastName  string   `xml:"lastname" json:"lastname,omitempty"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

//. Address
type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {
	rawXmlData := `
		<data>
			<person>
				<firstname>Nic</firstname>
				<lastname>Fury</lastname>
				<address>
					<city>Tracy</city>
					<state>CA</state>
				</address>
			</person>
			<person>
				<firstname>Maria</firstname>
				<lastname>Fury</lastname>
			</person>
		</data>
	`

	var data Data                            //=> initialize Users array
	xml.Unmarshal([]byte(rawXmlData), &data) //=> byteValue -> users
	jsonData, _ := json.Marshal(data)        //=> marshal to JSON
	fmt.Println(string(jsonData))            //=> print out
}
