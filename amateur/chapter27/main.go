package main

import "fmt"

//!เขียนโปรแกรมสี่เหลี่ยมจตุรัสพื้นหลังเลขเรียง 0 ถึง รับจากคีย์บอร์ด (ถ้าเลข>10 แสดงหลักหน่วยเท่านั้น)

func main() {
	n := 2
	pattern := (2 * n) + 1

	for i := 0; i < pattern; i++ {
		for j := 0; j < pattern; j++ {
			if i == 0 || i == pattern-1 || j == 0 || j == pattern-1 {
				fmt.Print(n - n)
			} else if i == (pattern-1)/2 && j == (pattern-1)/2 {
				fmt.Print(n)
			} else {
				fmt.Print(n - 1)
			}
		}
		fmt.Println()
	}

}
