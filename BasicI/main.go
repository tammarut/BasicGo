package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var the, day, gone int = 1, 2, 3

//████████████████████████████████████████████████████
func setadd(first, second int) int {
	return first + second
}

func swap(first, second string) (string, string) {
	return second, first
}

func split(para int) (first, second int) {
	first = para / 2
	second = para * 2
	return
}

func cut(first, second float32) (float32, error) {
	if second == 0.0 {
		err := errors.New("cut by zero!")
		return 0.0, err
	}
	return first / second, nil
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2)-x)/(2*z)
	}
	return z
}

//
func switchfall() {
	i := 1
	switch i {
	case 1:
		fmt.Print("one ")
		fallthrough
	case 2:
		fmt.Print("two ")
		fallthrough
	case 3:
		fmt.Print("three ")
		fallthrough
	default:
		fmt.Println("Go!")
	}
}

func switches() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}
func switchdate() {
	fmt.Println("When's Friday")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In 2 days.")
	default:
		fmt.Println("Too far away.")
	}
}
func switchNoCon() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}
}

//████████████████████████████████████████████████████
func main() {
	var naya = true
	box1 := "Arima1" //string เก็บ Arima1

	var arraybox [3]string
	arraybox[0] = "Dew"
	arraybox[1] = "Pleum"
	arraybox[2] = "Tankwa"

	arrayboxyao := [3]string{"Dew2", "Pleum2", "Tankwa3"}
	slicebox := []string{"Dew3", "Pleum3", "Tankwa3"}

	fmt.Println(box1)

	for i := 0; i < 3; i++ {
		fmt.Println(arraybox[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Println(arrayboxyao[i])
	}
	for i, list := range slicebox {
		fmt.Println(i, list)
	}
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(setadd(2, 3))
	fmt.Println(swap("Player1", "Player2"))
	fmt.Println(split(10))
	fmt.Println(the, day, gone, naya)

	result, err := cut(10, 2)
	if err != nil {
		//os.Exit
	}
	fmt.Println(result)
	const Pi = 3.14
	const World = "Yes or Yes"
	const Truth = true
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	var sums = 1
	for sums <= 10 { //while
		sums += sums
		sums++
	}
	fmt.Println(sums)
	fmt.Println(pow(2, 3, 10), pow(5, 2, 30))
	fmt.Println(sqrt(2))
	switches()
	switchdate()
	switchNoCon()
	switchfall()
	defer fmt.Print("Tammarut")
	fmt.Print("Pleum ")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}
}
