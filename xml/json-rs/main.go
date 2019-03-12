package main

//! Response JSON format by net/http
import (
	"encoding/json"
	"fmt"
	"net/http"
)

//. Profile ...
type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", foo)         //=> path "/"
	http.ListenAndServe(":3000", nil) //=> port, handler
}

func foo(w http.ResponseWriter, r *http.Request) { //=>handler
	w.Header().Set("Content-Type", "application/json") //=>set header

	profile := Profile{"Alex", []string{"snowboarding", "coding"}} //=> initialze data
	js, err := json.Marshal(profile)                               //=>marshal | Go>>JSON
	if err != nil {                                                //=>handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)             //=> response
	fmt.Println(string(js)) //=>print out
}
