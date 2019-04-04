package main

//! 3 ways; checking type variable
import (
	"fmt"
	"reflect"
)

func main() {
	//. Solution1: reflect
	floatna := 1.2
	fmt.Printf("floatna: %#[1]v  => %[1]T \n", reflect.TypeOf(floatna).String())

	if reflect.TypeOf(floatna).String() != "int" {
		fmt.Println("ไม่ใช่int")
	}

	fmt.Println("==================================")
	//. Solution2: func return Sprintf
	intkub := 1
	fmt.Println(func(v interface{}) string {
		return fmt.Sprintf("%T", v)
	}(intkub))
	if typeOf(intkub) != "float64" {
		fmt.Println("ไม่ใช่float")
	}

	fmt.Println("==================================")
	//. Solution3: assertion
	intka := 2
	var i interface{} = intka

	_, ok := i.(int)
	if ok {
		fmt.Println("นี่ไม่ใช่float64")
	}
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
