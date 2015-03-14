package main

import (
	"fmt"
)

// START OMIT
func do(f func(int) int, i int) int {
	return f(i)
}

func main() {
	plus1 := func(i int) int {
		return i + 1
	}
	fmt.Println(do(plus1, 7))
}

//END OMIT
