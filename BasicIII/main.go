package main

import "fmt"

var first, second int

func plus(first, second int) int {
	return first + second
}
func prism(length, width, height int) int {
	return length * width * height
}

func first_and_last(mac string) (string, string) {
	return string(mac[0]), string(mac[len(mac)-6])
}
func swap(first, second int) (int, int) {
	return second, first
}
func sums(number ...int) { //Variadic function
	fmt.Print(number, " = ")

	total := 0
	for _, i := range number {
		total += i
	}
	fmt.Println(total)
}
func seal() func() int { //<--Closesure
	i := 0
	return func() int {
		i++
		return i
	}
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func factForloop(n int) int {
	total := 1
	for i := 1; i <= 5; i++ {
		total = total * i
	}
	return total
}

func main() {
	// fmt.Print("Input first: ")
	// fmt.Scan(&first)
	// fmt.Print("Input second: ")
	// fmt.Scan(&second)

	// result := plus(first, second)
	// fmt.Println("plus= ", result)

	// result = prism(first, second, 2)
	// fmt.Println("prism= ", result)

	// thrid, fouth := first_and_last("Hello Kaneki")
	// fmt.Println(thrid + " " + fouth)

	// fifth, sixth := swap(0, 1)
	// fmt.Println(fifth, sixth)

	// sums(1, 2) //<<--Variadic function
	// sums(1, 2, 3)
	// number := []int{10, 20, 30, 40, 50}
	// sums(number[0:4]...)
	// sums(number...)

	// nextInt := seal() //Closures
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// fmt.Println("Factorial Recursion= ", fact(5))
	// fmt.Println("Factorial Forloop= ", factForloop(5))
}

func init() {
	fmt.Println("Init first..\n---------------------------")
}
