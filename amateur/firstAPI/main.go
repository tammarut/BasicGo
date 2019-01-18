package main

//! Use Gorillamux => It's work Ok!
import (
	"fmt"
	"net/http"

	mux "mod/github.com/gorilla/mux@v1.6.2"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/greeting/{name}", greeting)

	fmt.Println("Start..")
	http.ListenAndServe(":1000", mux)
}

func homepage(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Header().Set("Contenttype", "application")
	fmt.Fprintf(w, `{"message":"welcome to homepage"}`)
}

func greeting(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	w.Header().Set("Contenttype", "application")
	fmt.Fprintf(w, `{"message":"welcome to homepage"}`)
	fmt.Fprintf(w, `{"message":"hi arima %s"}`, name)
}
