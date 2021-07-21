package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	x := 2*b - a - c
	var k int
	if x >= 0 {
		k = 0
	} else {
		k = (1 - x) / 2
	}
	fmt.Println(x + 3*k)
}
