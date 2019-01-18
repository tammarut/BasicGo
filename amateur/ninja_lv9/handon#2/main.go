package main

import "fmt"

//! Method set [struct, interface, method]
type person struct {
	name string
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("You're great.")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	player1 := person{"kaneki"}

	saySomething(&player1) //TODO: same result
	fmt.Println("————————————————————————————————————————")
	player1.speak() //TODO: same result

}
