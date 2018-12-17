package main

import (
	"fmt"
)

// func setZeroPassByVal(i int) {
// 	i = 0
// 	fmt.Println("After set PassbyValue: ", i)
// }
// func setZeroPassByRef(i *int) {
// 	*i = 0
// 	fmt.Println("After set PassbyRef: ", *i)
// }
func main() {
	// fmt.Println("Hi, defer")
	// defer fmt.Print("one ")
	// defer fmt.Print("two ")
	// defer fmt.Print("three ")

	/*	defer func() { //Recover
			if recov := recover(); recov != nil {
				fmt.Println("Recovered: The system can't find path specified.")
			}
		}()

		//panic below
		path := "/Users/Admin/Goo/src/github.com/khup/BasicIV/password.txt"
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			file, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}
		fmt.Println("Done creating file.", path) */

	// di := 1
	// fmt.Println("Value di: ", di)  //data
	// fmt.Println("Index di: ", &di) //index

	// fmt.Println("--------------------------------")

	// p := &di
	// fmt.Println("Value p point: ", p)
	// fmt.Println("Index p: ", &p)
	// fmt.Println("Value *p: ", *p)
	// fmt.Println("Size p: ", unsafe.Sizeof(p), "bytes")

	// fmt.Println("--------------------------------")

	// *p = 10
	// fmt.Println("Value *p: ", *p)
	// fmt.Println("Value i: ", di)

	// i := 5
	// fmt.Println("Value i: ", i)
	// setZeroPassByVal(i)
	// fmt.Println("Value i after: ", i)

	// fmt.Println("Value i: ", i)
	// setZeroPassByRef(&i)
	// fmt.Println("Value i after: ", i)
}
func init() {
	fmt.Println("Run Init first..")
	fmt.Println("---------------------------")
}
