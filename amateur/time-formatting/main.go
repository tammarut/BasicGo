package main

//! Time formatting/Parsing
import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

	p("1)", t1) //=> 2012-11-01 22:08:41 +0000 +0000

	p(t.Format("3:04PM"))                          //=> 11.48PM
	p(t.Format("Mon Jan _2 15:05:05 2006"))        //=> Wed Feb  6 23:58:58 2019
	p(t.Format("2006-01-02T15:04:05.99999-07:00")) //=>2019-02-06T23:48:58.66365+07:00

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2) //=> 0000-01-01 20:41:00 +0000 UTC

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) //=> 2019-02-06T23:48:58-00:00

	errsure := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(errsure, "8:41PM") //!Error parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
	p(e)
}
