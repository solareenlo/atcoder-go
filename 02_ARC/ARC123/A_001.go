package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	x := 2*b - a - c
	k := max(0, (-x+1)/2)
	fmt.Println(x + 3*k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
