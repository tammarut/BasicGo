package fizzbuzz

import "strconv"

func Fizzbuzz(number int) string {
	if number == 3 {
		return "Fizz"
	} else if number == 5 {
		return "Buzz"
	} else if number == 15 {
		return "FizzBuzz"
	}
	return strconv.Itoa(number)
}
