package main

import "fmt"

//! handling_Error(panic) and input from keyboard(Scan)
func main() {
	var name, kind, time string

	fmt.Print("Name:")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err) //TODO:panic right here!!
	}

	fmt.Print("Kind:")
	_, err = fmt.Scan(&kind)
	if err != nil {
		panic(err) //TODO:panic right here!!
	}

	fmt.Print("Time:")
	_, err = fmt.Scan(&time)
	if err != nil {
		panic(err) //TODO:panic right here!!
	}

	fmt.Println(name, kind, time)
}

func init() {
	fmt.Println("Hi, handle_error 02")
	fmt.Println("---------------------")
}
