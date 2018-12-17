package main

import (
	"fmt"
)

type Vertex struct {
	first, seconcd int
}

var (
	vertex1 = Vertex{1, 2}
	vertex2 = Vertex{first: 1}
	vertex3 = Vertex{}
	p       = &Vertex{1, 2}
)

//████████████████████████████████████████████████████
func pointer() {
	i, j := 40, 100

	p := &i         // &reference
	fmt.Println(*p) // *dereference

	*p = 1
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)
}
func slices_from_array() {
	arraybox := [5]int{1, 2, 3, 4, 5}

	var asd []int = arraybox[0:2]
	fmt.Println(asd)
}
func sliceRefArr() {
	names := [4]string{"Joh", "Pao", "Seek", "Cha"}
	fmt.Println(names, " <<total member's array")

	a := names[0:2]
	b := names[2:]
	fmt.Println(a, b, " <<2 slices cut ")

	b[1] = "Arima"
	fmt.Println(a, b, " <<add Arima b[1]")
	fmt.Println(names, " <<total member's array after adding")
}
func copyAppend() {
	first := []string{"Alex", "Jap"}
	second := make([]string, len(first))

	copy(second, first)
	second = append(second, "Ducati", "Honda")
	fmt.Printf("first slice %d: %v\n", len(first), first)
	fmt.Printf("second slice %d: %v \n", len(second), second)
}
func slicePoint() {
	first := []string{"zero", "one", "two", "three", "four"}
	second := first[0:5]
	third := first[:5]
	fourth := first[0:]
	fifth := first[:]
	sixth := fifth[1:3]

	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println(sixth, len(sixth), cap(sixth))
}

func slicenil() {
	var sn []int
	fmt.Println(sn, len(sn), cap(sn))
	if sn == nil {
		fmt.Println("It's nill slice.")
	}
}
func slicemake() {
	a := make([]int, 5)
	b := make([]int, 0, 5)
	c := b[:2]
	d := c[2:5]
	printSlice("a", a)
	printSlice("b", b)
	printSlice("c", c)
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//████████████████████████████████████████████████████
/*func main() {
	//pointer()
	// fmt.Println(Vertex{1, 1}) //Vertex = type struct
	//v := Vertex{1, 1}
	// v.first = 0
	// v.seconcd = 0
	// fmt.Println(v.first, v.seconcd, " ", Vertex{1, 1})

	// p := &v
	// p.first = 1e9
	// fmt.Println(p)
	// fmt.Println(vertex1, vertex2, vertex3, p)
	// fmt.Println("_________________________________")
	//slices_from_array()
	//sliceRefArr()

	//slicenil()
	//slicemake()
	//copyAppend()
	//slicePoint()
	// fmt.Println("_________________________________")
	//var keyboard int
	//fmt.Print("Input: ")
	//fmt.Scan(&keyboard)
	//fmt.Printf("Keyboard= %d", keyboard)
} */
func init() {
	fmt.Println("Run Init first.\n__________________________")
}
