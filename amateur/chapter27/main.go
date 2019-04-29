package main

import (
	"fmt"
	"math"
)

//!เขียนโปรแกรมสี่เหลี่ยมจตุรัสพื้นหลังเลขเรียง 0 ถึง รับจากคีย์บอร์ด (ถ้าเลข>10 แสดงหลักหน่วยเท่านั้น)

func main() {
	var absi, absj float64
	n := 13
	n = n % 10

	//. Solution2
	for i := 0; i <= 2*n; i++ {
		for j := 0; j <= 2*n; j++ {
			absi = float64(n) - math.Abs((float64(n - i)))
			absj = float64(n) - math.Abs((float64(n - j)))

			if absi < absj {
				fmt.Print(absi)
			} else {
				fmt.Print(absj)
			}
		}
		fmt.Println()
	}

	//. Solution1 >>Hard code
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 || i == n-1 || j == 0 || j == n-1 {
	// 			fmt.Print(n - n)
	// 		} else if i == (n-1)/2 && j == (n-1)/2 {
	// 			fmt.Print(n)
	// 		} else {
	// 			fmt.Print(n - 1)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}
