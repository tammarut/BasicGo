package main

//! Method pointer
import (
	"fmt"
	"math"
)

//Paper Normal struct
type Paper struct {
	X float64
	Y float64
}

//Scale is Paper's method
func (v *Paper) Scale(f float64) { //=>ล้วงเข้าไปแก้ค่าตัวเดียวกันได้
	v.X *= f //=> change value on same fish
	v.Y *= f //=> change value on same fish
}

//Abs ...
func (v Paper) Abs() float64 { //=> method attach struct, it's easy to call
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	a4 := Paper{3, 4}
	a4.Scale(10)          //=> Call method
	fmt.Println(a4.Abs()) //=>Call and print out
}
