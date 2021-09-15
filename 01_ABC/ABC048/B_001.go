package main

import "fmt"

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	cntA := -1
	if a != 0 {
		cntA = (a - 1) / x
	}
	cntB := b / x
	fmt.Println(cntB - cntA)
}
