package main

import "github.com/tammarut/basicGo/amateur/echoYoutube/router"

//! Golang API(echo) =>> learn from Youtube

func main() {
	e := router.New()
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
