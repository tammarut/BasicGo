package main

//! Response XML by net/http
import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/", foo)         //=> path "/"
	http.ListenAndServe(":3000", nil) //=> port 3000
}
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")              //=> set Header
	profile := Profile{"Alex", []string{"snowboarding", "coding"}} //=> intialize data

	x, err := xml.MarshalIndent(profile, "", " ") //=> like mashal as new indent
	if err != nil {                               //=> handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(x)             //=> response
	fmt.Println(string(x)) //=> print log out
}

//curl -i localhost:3000
/*
HTTP/1.1 200 OK
Content-Type: application/xml
Content-Length: 128

<Profile>
  <Name>Alex</Name>
  <Hobbies>
    <Hobby>snowboarding</Hobby>
    <Hobby>coding</Hobby>
  </Hobbies>
</Profile>
*/
