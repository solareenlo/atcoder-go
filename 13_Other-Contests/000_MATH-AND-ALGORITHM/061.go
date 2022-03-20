package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var a int
	fmt.Scan(&a)
	a++

	if bits.OnesCount(uint(a)) == 1 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
