package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	f := a + b
	fmt.Println(max(max(a+b+c, a*b*c), max(f*c, a*b+c)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
