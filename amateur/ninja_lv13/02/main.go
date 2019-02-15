package main

//! main call func word.Count and forloop func word.UseCount
//! Nothing much just call func.
import (
	"fmt"

	"github.com/tammarut/amateur/ninja_lv13/02/quote"
	"github.com/tammarut/amateur/ninja_lv13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(k, v)
	}
}
