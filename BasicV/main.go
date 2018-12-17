package main

import (
	"errors"
	"fmt"
	"math"
)

type developer struct {
	skill    []string
	exp      int
	salary   int
	location string
	line     string
	company  string
}

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (point rect) area() float64 {

	return point.width * point.height
}
func (point rect) perim() float64 {

	return 2*point.width + 2*point.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perim: ", g.perim())
}
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("You can't divide by zero.")
	}
	return x / y, nil
}
func main() {
	// pleum := developer{
	// 	skill:    []string{"java-", "c#"},
	// 	exp:      0,
	// 	location: "nonthaburi",
	// 	line:     "n/a",
	// 	company:  "Odd-e",
	// }
	// pleum.salary = 30000
	// pleum.exp = 2
	// fmt.Println(pleum)
	// fmt.Println("---------------------------")

	// dew := &pleum
	// fmt.Println(dew.company)
	// dew.company = "Intervision"
	// fmt.Println(dew.company)
	// fmt.Println(pleum)

	// r := rect{width: 10, height: 2}
	// fmt.Println("area: ", r.area())
	// fmt.Println("perim: ", r.perim())

	// rp := &r
	// fmt.Println("area rp: ", rp.area())
	// fmt.Println("perim rp: ", rp.perim())

	// r := rect{width: 10, height: 2}
	// c := circle{radius: 3}

	// measure(r)
	// measure(c)

	if r, e := divide(2, -1); e != nil {
		fmt.Println("error: ", e)
	} else {
		fmt.Println("result: ", r)
	}
}

func init() {
	fmt.Println("Run init first..")
	fmt.Println("--------------------------------")
}
